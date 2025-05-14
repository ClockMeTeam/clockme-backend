package http

import (
	"github.com/maevlava/ftf-clockify/internal/app"
	"net/http"
)

func NewRouter(app *app.Application) http.Handler {
	mux := http.NewServeMux()

	return mux
}
