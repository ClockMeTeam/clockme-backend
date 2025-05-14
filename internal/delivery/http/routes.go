package http

import (
	"github.com/maevlava/ftf-clockify/internal/app"
	"net/http"
)

func NewRouter(app *app.Application, workDebtHandler *WorkDebtHandler) http.Handler {
	mux := http.NewServeMux()
	apiMux := apiServerMux(workDebtHandler)

	mux.Handle("/api/v1/", apiMux)
	return mux
}

func apiServerMux(workDebtHandler *WorkDebtHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /debts", workDebtHandler.GetAllUserWorkDebt)

	return mux
}
