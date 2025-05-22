package workdebt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maevlava/ftf-clockify/internal/config"
	"github.com/maevlava/ftf-clockify/internal/domain"
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

type WorkDebtResponse struct {
	Maevlava string `json:"maevlava"`
	Deandra  string `json:"deandra"`
}

type WorkDebtService interface {
	GetAllUserWorkDebt() (WorkDebtResponse, error)
	GetWorkDebtByProject(projectId string) (string, error)
}

type workDebtService struct {
	config   *config.ApiConfig
	userRepo domain.UserRepository
	//	projectRepo domain.ProjectRepository
}

func NewService(cfg *config.ApiConfig, userRepo domain.UserRepository) WorkDebtService {
	return &workDebtService{
		config:   cfg,
		userRepo: userRepo,
	}
}

// TODO refactor to database
func (w workDebtService) GetAllUserWorkDebt() (WorkDebtResponse, error) {
	users, err := w.userRepo.GetUsers(context.TODO())
	if err != nil {
		return WorkDebtResponse{}, err
	}
	for _, user := range users {
		log.Printf("User: %s\n", user.Name)
	}
	//owed, _ := calculateGrossWorkingHoursOwed()
	//actualMaevlavaWorkingHours, err := calculateTotalActualWorkingHours(w.config.WorkspaceId, w.config.MaevlavaId, w.config.ClockifySecret)
	//if err != nil {
	//	log.Printf(err.Error())
	//}
	//actualDeandraWorkingHours, err := calculateTotalActualWorkingHours(w.config.WorkspaceId, w.config.DeandraId, w.config.ClockifySecret)
	//if err != nil {
	//	log.Printf(err.Error())
	//}
	//
	//totalMaevlavaDebt := owed - actualMaevlavaWorkingHours
	//totalDeandraDebt := owed - actualDeandraWorkingHours
	//
	//return WorkDebtResponse{
	//	Maevlava: totalMaevlavaDebt.String(),
	//	Deandra:  totalDeandraDebt.String(),
	//}, nil
	return WorkDebtResponse{}, nil
}
func (w workDebtService) GetWorkDebtByProject(projectId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func calculateGrossWorkingHoursOwed() (time.Duration, error) {
	var grossWorkHoursOwed time.Duration

	baseWorkingHours := time.Duration(BASE_WORKING_HOURS) * time.Hour

	dateLayout := "2006-01-02"
	todayString := time.Now().Format(dateLayout)
	endDateLoop, err := time.ParseInLocation(dateLayout, todayString, time.UTC)
	if err != nil {
		return 0, err
	}
	startDateLoop, err := time.ParseInLocation(dateLayout, PROJECT_START, time.UTC)
	if err != nil {
		return 0, err
	}

	log.Printf("Iterating from %s to %s\n", startDateLoop, endDateLoop)
	currentDate := startDateLoop
	for !currentDate.After(endDateLoop) {
		dayOfWeek := currentDate.Weekday()

		if slices.Contains(HOLIDAYS, dayOfWeek.String()) {
			log.Printf("Skipping holiday %s\n", dayOfWeek.String())
			currentDate = currentDate.AddDate(0, 0, 1)
			continue
		}
		log.Printf("Processing day %s %s\n", dayOfWeek.String(), currentDate.String())

		grossWorkHoursOwed += baseWorkingHours
		log.Printf("Work Hours Owed %s", grossWorkHoursOwed.String())
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return grossWorkHoursOwed, nil
}
func calculateTotalActualWorkingHours(workspaceId, userId, apiKey string) (time.Duration, error) {
	log.Printf("CALCULATE TOTAL WORKING HOURS\nWORKSPACE: %s\nUSER: %s\nAPI_KEY: %s\n", workspaceId, userId, apiKey)
	type TimeInterval struct {
		Start string  `json:"start"`
		End   *string `json:"end"`
	}
	type Response struct {
		TimeInterval TimeInterval `json:"timeInterval"`
	}
	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.clockify.me/api/v1/workspaces/%s/user/%s/time-entries", workspaceId, userId),
		nil,
	)
	if err != nil {
		return 0, errors.New("failed to create request")
	}
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

	var totalDuration time.Duration
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
