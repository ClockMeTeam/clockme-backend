package http

import (
	"github.com/clockme/clockme-backend/internal/app"
	"net/http"
)

func NewRouter(app *app.Application, workDebtHandler *WorkDebtHandler) http.Handler {
	mux := http.NewServeMux()
	apiMux := apiServerMux(workDebtHandler)

	apiHandler := http.StripPrefix("/api", apiMux)

	mux.Handle("/api/", apiHandler)
	return mux
}

func apiServerMux(workDebtHandler *WorkDebtHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/debts", workDebtHandler.GetUsersWorkDebt)
	mux.HandleFunc("GET /users/project-debts", workDebtHandler.GetUsersWorkDebtByType)

	return mux
}
