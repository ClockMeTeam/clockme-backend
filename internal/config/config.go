package config

import (
	"os"
)

type ApiConfig struct {
	ClockifySecret string
	MaevlavaId     string
	DeandraId      string
	WorkspaceId    string
}
type DatabaseConfig struct {
	Name     string
	Host     string
	Password string
	Port     string
	User     string
	SSLMode  string
}
type AppConfig struct {
	InternalPort string
}
type Config struct {
	App      AppConfig
	API      ApiConfig
	Database DatabaseConfig
}

func Load() *Config {
	return &Config{
		App:      loadAppConfig(),
		API:      loadApiConfig(),
		Database: loadDatabaseConfig(),
	}
}
func loadAppConfig() AppConfig {
	internalPort := os.Getenv("APP_INTERNAL_PORT")
	return AppConfig{
		InternalPort: internalPort,
	}
}
func loadDatabaseConfig() DatabaseConfig {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	return DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
		User:     dbUser,
		Password: dbPassword,
		SSLMode:  dbSSLMode,
	}
}
func loadApiConfig() ApiConfig {
	clockifySecret := os.Getenv("CLOCKIFY_API_KEY")
	MaevlavaId := os.Getenv("MAEVLAVA_ID")
	DeandraId := os.Getenv("DEANDRA_ID")
	WorkspaceId := os.Getenv("WORKSPACE_ID")

	return ApiConfig{
		ClockifySecret: clockifySecret,
		MaevlavaId:     MaevlavaId,
		DeandraId:      DeandraId,
		WorkspaceId:    WorkspaceId,
	}
}
