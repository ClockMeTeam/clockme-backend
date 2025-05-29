package main

import (
	"fmt"
	"github.com/maevlava/ftf-clockify/internal/repository"

	//"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/maevlava/ftf-clockify/internal/app"
	"github.com/maevlava/ftf-clockify/internal/config"
	httpdelivery "github.com/maevlava/ftf-clockify/internal/delivery/http"
	"github.com/maevlava/ftf-clockify/internal/repository/postgres"
	"github.com/maevlava/ftf-clockify/internal/service/workdebt"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv()

	// app
	cfg := config.Load()
	appInstance := app.NewApp(cfg)

	// database
	dbPool, err := postgres.NewConnectionPool(cfg.Database)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error connecting to database: %s", err))
	}
	defer dbPool.Close()

	// repositories
	userRepo := repository.NewPgUserRepository(dbPool)
	projectRepo := repository.NewPgProjectRepository(dbPool)
	projectTypeRepo := repository.NewPgProjectTypeRepository(dbPool)

	// Services
	workDebtService := workdebt.NewService(&cfg.API, userRepo, projectRepo, projectTypeRepo)

	// Handlers
	workDebtHandler := httpdelivery.NewWorkDebtHandler(workDebtService)

	port := os.Getenv("APP_INTERNAL_PORT")
	address := ":" + port
	router := httpdelivery.NewRouter(appInstance, workDebtHandler)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(server.ListenAndServe())
}

func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		errLoad := godotenv.Load()
		if errLoad != nil {
			log.Println("Warning: .env file found but could not be loaded:", errLoad)
		} else {
			log.Println("Loaded environment variables from .env file")
		}
	} else {
		log.Println("No .env file found, relying on OS environment variables.")
	}
}
