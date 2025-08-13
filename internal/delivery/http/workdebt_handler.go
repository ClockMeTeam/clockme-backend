package http

import (
	"github.com/clockme/clockme-backend/internal/service/workdebt"
	"log"
	"net/http"
)

type WorkDebtHandler struct {
	service workdebt.WorkDebtService
}

func NewWorkDebtHandler(service workdebt.WorkDebtService) *WorkDebtHandler {
	return &WorkDebtHandler{
		service: service,
	}
}

func (h *WorkDebtHandler) GetUsersWorkDebt(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		HoursOwed string `json:"hours_owed"`
	}
	var response []Response

	users, err := h.service.GetUsersWorkDebt()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	for _, user := range users {
		response = append(response, Response{
			Name:      user.Name,
			Email:     user.Email,
			HoursOwed: user.HoursOwed.String(),
		})
	}

	RespondWithJSON(w, http.StatusOK, response)
}
func (h *WorkDebtHandler) GetUsersWorkDebtByType(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Name      string            `json:"name"`
		Email     string            `json:"email"`
		HoursOwed map[string]string `json:"hours_owed"`
	}
	var response []Response
	users, usersTypeHours, err := h.service.GetWorkDebtByProjectType()
	if err != nil {
		log.Printf("Error getting work debt by project type: %v", err)
	}
	for i, user := range users {
		response = append(response, Response{
			Name:      user.Name,
			Email:     user.Email,
			HoursOwed: usersTypeHours[i],
		})
	}
	RespondWithJSON(w, http.StatusOK, response)
}
