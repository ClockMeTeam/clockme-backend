package config

import "os"

type ApiConfig struct {
	ClockifySecret string
}

func Load() *ApiConfig {
	clockifySecret := os.Getenv("CLOCKIFY_SECRET")
	return &ApiConfig{
		ClockifySecret: clockifySecret,
	}
}
