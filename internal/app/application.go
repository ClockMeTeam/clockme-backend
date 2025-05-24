package app

import "github.com/maevlava/ftf-clockify/internal/config"

type Application struct {
	Config *config.Config
}

func NewApp(cfg *config.Config) *Application {
	return &Application{
		Config: cfg,
	}
}
