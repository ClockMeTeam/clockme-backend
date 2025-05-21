package app

import "github.com/maevlava/ftf-clockify/internal/config"

type Application struct {
	Config *config.ApiConfig
}

func NewApp(cfg *config.ApiConfig) *Application {
	return &Application{
		Config: cfg,
	}
}
