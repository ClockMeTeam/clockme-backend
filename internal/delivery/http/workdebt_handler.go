package http

import (
	"github.com/maevlava/ftf-clockify/internal/service/workdebt"
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

func (h *WorkDebtHandler) GetAllUserWorkDebt(w http.ResponseWriter, r *http.Request) {
	workDebtResponse, err := h.service.GetAllUserWorkDebt()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, workDebtResponse)
}
