package app

import "github.com/maevlava/ftf-clockify/internal/config"

type App struct {
	Config *config.ApiConfig
}

func NewApp() *App {
	return &App{}
}
