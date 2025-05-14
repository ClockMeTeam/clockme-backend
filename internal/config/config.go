package config

import "os"

type ApiConfig struct {
	ClockifySecret string
}

func Load() *ApiConfig {
	clockifySecret := os.Getenv("CLOCKIFY_API_KEY")
	return &ApiConfig{
		ClockifySecret: clockifySecret,
	}
}
