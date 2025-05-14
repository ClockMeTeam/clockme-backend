package config

import "os"

type ApiConfig struct {
	ClockifySecret string
	MaevlavaId     string
	DeandraId      string
	WorkspaceId    string
}

func Load() *ApiConfig {
	clockifySecret := os.Getenv("CLOCKIFY_API_KEY")
	MaevlavaId := os.Getenv("MAEVLAVA_ID")
	DeandraId := os.Getenv("DEANDRA_ID")
	ProjectId := os.Getenv("PROJECT_ID")
	return &ApiConfig{
		ClockifySecret: clockifySecret,
		MaevlavaId:     MaevlavaId,
		DeandraId:      DeandraId,
		WorkspaceId:    ProjectId,
	}
}
