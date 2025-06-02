package workdebt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clockme/clockme-backend/internal/config"
	"github.com/clockme/clockme-backend/internal/domain"
	"io"
	"log"
	"net/http"
	"slices"
	"time"
)

/* List Api Clockify
* GetAllWorkspace	https://api.clockify.me/api/v1/workspaces
* GetALlUsers		https://api.clockify.me/api/v1/workspaces/{workspaceId}/users
* GetAllProjects	https://api.clockify.me/api/v1/workspaces/{workspaceId}/projects
* GetProjectById	https://api.clockify.me/api/v1/workspaces/{workspaceId}/projects/{projectId}
* Time Entries		https://api.clockify.me/api/v1/workspaces/{workspaceId}/user/{userId}/time-entries
 */

const BASE_WORKING_HOURS = 6
const PROJECT_START = "2025-05-09"

var HOLIDAYS = []string{"Tuesday", "Thursday"}

type WorkDebtService interface {
	GetUsersWorkDebt() ([]domain.User, error)
	GetWorkDebtByProjectType() ([]domain.User, []map[string]string, error)
}

type workDebtService struct {
	config          *config.ApiConfig
	userRepo        domain.UserRepository
	projectRepo     domain.ProjectRepository
	projectTypeRepo domain.ProjectTypeRepository
}

func NewService(
	cfg *config.ApiConfig,
	userRepo domain.UserRepository,
	projectRepo domain.ProjectRepository,
	projectTypeRepo domain.ProjectTypeRepository) WorkDebtService {

	return &workDebtService{
		config:          cfg,
		userRepo:        userRepo,
		projectRepo:     projectRepo,
		projectTypeRepo: projectTypeRepo,
	}
}

func (w workDebtService) GetUsersWorkDebt() ([]domain.User, error) {
	users, err := w.userRepo.GetUsers(context.TODO())
	if err != nil {
		return []domain.User{}, err
	}

	owe, err := calculateGrossWorkingHoursOwed()
	if err != nil {
		return []domain.User{}, err
	}

	for i, user := range users {
		actual, err := calculateTotalActualWorkingHours(w.config.WorkspaceId, user.ClockifyID, w.config.ClockifySecret)
		if err != nil {
			return []domain.User{}, err
		}
		users[i].HoursOwed = owe - actual
	}

	return users, nil
}
func (w workDebtService) GetWorkDebtByProjectType() ([]domain.User, []map[string]string, error) {
	// Get all project types
	projectTypes, err := w.projectTypeRepo.GetProjectTypes(context.TODO())
	if err != nil {
		return nil, nil, fmt.Errorf("error getting project type: %w", err)
	}
	// Get gross base hours per type
	grossHoursPerType, err := calculateGrossWorkingHoursOwedByProjectType(w.projectTypeRepo)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting project type: %w", err)
	}
	// get all users
	users, err := w.userRepo.GetUsers(context.TODO())
	if err != nil {
		return nil, nil, err
	}

	var readableUserTypeDebts []map[string]string
	for _, user := range users {
		userWorkedHours, err := filterProjectsFromTimeEntries(w.config.WorkspaceId, user.ClockifyID, w.config.ClockifySecret, projectTypes, w.projectRepo)
		if err != nil {
			return nil, nil, err
		}
		userDebt := make(map[string]string)
		for projectType, gross := range grossHoursPerType {
			actual := userWorkedHours[projectType]
			debt := gross - actual
			log.Printf("debt for %s: %s\n", projectType, debt.String())
			userDebt[projectType] = debt.String()
		}
		readableUserTypeDebts = append(readableUserTypeDebts, userDebt)
	}

	return users, readableUserTypeDebts, nil
}

// Improve to reduce redundant db calls on the same projectId (now using cache)
// TODO improve using go routines
func filterProjectsFromTimeEntries(workspaceId, userId, apiKey string, projectTypes []domain.ProjectType, projectRepo domain.ProjectRepository) (map[string]time.Duration, error) {
	// get user time entries
	type Response struct {
		TimeInterval TimeInterval `json:"timeInterval"`
		ProjectId    string       `json:"projectId"`
	}

	// variable total hours per types
	typeHours := make(map[string]time.Duration)

	page := 1
	pageSize := 100
	for {
		// request
		resp, err := requestGetTimeEntries(workspaceId, userId, apiKey, &page, &pageSize)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		var response []Response
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		if len(response) == 0 {
			break
		}
		// get the intervals and the projectIds
		projectTypeCache := make(map[string]domain.ProjectType)
		for _, entry := range response {
			if entry.TimeInterval.End == nil {
				continue
			}

			pt, ok := projectTypeCache[entry.ProjectId]
			if !ok {
				projectType, err := projectRepo.GetProjectTypeByClockifyID(context.TODO(), entry.ProjectId)
				if err != nil {
					log.Printf("Missing project in DB for clockify_id: %s", entry.ProjectId)
					return nil, fmt.Errorf("error getting project type: %w", err)
				}
				if projectType.Name == "" {
					log.Printf("Project found but has no type assigned (empty name). Clockify projectId: %s", entry.ProjectId)
					continue // or assign projectType.Name = "Unknown"
				}
				projectTypeCache[entry.ProjectId] = projectType
				pt = projectType
			}

			duration, err := calculateTimeInterval(entry.TimeInterval.Start, *entry.TimeInterval.End)
			if err != nil {
				return nil, fmt.Errorf("failed to calculate time interval err: %w", err)
			}

			// add type hours for type in projectId
			typeHours[pt.Name] += duration
		}
		page++
	}
	return typeHours, nil
}
func requestGetTimeEntries(workspaceId, userId, apiKey string, page, pageSize *int) (*http.Response, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://api.clockify.me/api/v1/workspaces/%s/user/%s/time-entries?page=%d&page-size=%d", workspaceId, userId, *page, *pageSize)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request")
	}
	req.Header.Add("X-Api-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do time entries request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("server error: status code %d, body: %s", resp.StatusCode, string(body))
	}
	return resp, nil
}
func calculateGrossWorkingHoursOwed() (time.Duration, error) {
	var grossWorkHoursOwed time.Duration

	baseWorkingHours := time.Duration(BASE_WORKING_HOURS) * time.Hour

	dateLayout := "2006-01-02"
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	todayString := time.Now().In(jakarta).Format(dateLayout)
	endDateLoop, err := time.ParseInLocation(dateLayout, todayString, jakarta)
	if err != nil {
		return 0, errors.New("failed to parse today's date")
	}
	startDateLoop, err := time.ParseInLocation(dateLayout, PROJECT_START, jakarta)
	if err != nil {
		return 0, errors.New("failed to parse today's date")
	}

	log.Printf("Iterating from %s to %s\n", startDateLoop, endDateLoop)
	currentDate := startDateLoop
	for !currentDate.After(endDateLoop) {
		dayOfWeek := currentDate.Weekday()

		if slices.Contains(HOLIDAYS, dayOfWeek.String()) {
			currentDate = currentDate.AddDate(0, 0, 1)
			continue
		}

		grossWorkHoursOwed += baseWorkingHours
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return grossWorkHoursOwed, nil
}
func calculateGrossWorkingHoursOwedByProjectType(projectTypeRepo domain.ProjectTypeRepository) (map[string]time.Duration, error) {
	projectTypes, err := projectTypeRepo.GetProjectTypes(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error getting project type: %w", err)
	}

	baseHoursPerType := make(map[string]time.Duration)
	for _, projectType := range projectTypes {
		log.Printf("Project Hour: %d\n", projectType.BaseHour)
		baseHoursPerType[projectType.Name] = time.Duration(projectType.BaseHour) * time.Hour
		log.Printf("Base hours for %s: %s\n", projectType.Name, baseHoursPerType[projectType.Name])
	}

	dateLayout := "2006-01-02"
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	todayString := time.Now().In(jakarta).Format(dateLayout)
	endDateLoop, err := time.ParseInLocation(dateLayout, todayString, jakarta)
	if err != nil {
		return nil, errors.New("failed to parse today's date")
	}
	startDateLoop, err := time.ParseInLocation(dateLayout, PROJECT_START, jakarta)
	if err != nil {
		return nil, errors.New("failed to parse today's date")
	}
	log.Printf("Iterating from %s to %s\n", startDateLoop, endDateLoop)

	totalHoursPerType := make(map[string]time.Duration)
	currentDate := startDateLoop
	for !currentDate.After(endDateLoop) {
		dayOfWeek := currentDate.Weekday().String()

		if slices.Contains(HOLIDAYS, dayOfWeek) {
			currentDate = currentDate.AddDate(0, 0, 1)
			continue
		}

		for name, baseHours := range baseHoursPerType {
			totalHoursPerType[name] += baseHours
			//log.Printf("Total hours for %s: %s\n", name, totalHoursPerType[name])
		}

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return totalHoursPerType, nil
}
func calculateTotalActualWorkingHours(workspaceId, userId, apiKey string) (time.Duration, error) {
	log.Printf("CALCULATE TOTAL WORKING HOURS\nWORKSPACE: %s\nUSER: %s\nAPI_KEY: %s\n", workspaceId, userId, apiKey)
	type Response struct {
		TimeInterval TimeInterval `json:"timeInterval"`
	}
	client := &http.Client{}
	page := 1
	pageSize := 100
	var totalDuration time.Duration

	for {
		// request
		url := fmt.Sprintf("https://api.clockify.me/api/v1/workspaces/%s/user/%s/time-entries?page=%d&page-size=%d", workspaceId, userId, page, pageSize)
		req, err := http.NewRequest(
			http.MethodGet,
			url,
			nil,
		)
		if err != nil {
			return 0, errors.New("failed to create request")
		}

		// Header with Clockify-key
		req.Header.Add("X-Api-Key", apiKey)

		// test ca certificate please
		resp, err := client.Do(req)
		if err != nil {
			return 0, errors.New(fmt.Sprintf("failed to get work debt from server\nerror: %v", err))
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return 0, errors.New("failed to get Maevlava work debt")
		}

		var response []Response

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return 0, errors.New("failed to decode response")
		}

		// break if response is 0 (nothing in page)
		if len(response) == 0 {
			break
		}

		for _, entry := range response {

			if entry.TimeInterval.End == nil {
				continue
			}

			duration, err := calculateTimeInterval(entry.TimeInterval.Start, *entry.TimeInterval.End)
			if err != nil {
				return 0, errors.New("failed to calculate time interval")
			}
			totalDuration += duration
		}

		page++
	}
	return totalDuration, nil
}
func calculateTimeInterval(start, end string) (time.Duration, error) {
	sTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return 0, err
	}
	eTime, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return 0, err
	}
	duration := eTime.Sub(sTime)

	return duration, nil
}

type TimeInterval struct {
	Start string  `json:"start"`
	End   *string `json:"end"`
}
