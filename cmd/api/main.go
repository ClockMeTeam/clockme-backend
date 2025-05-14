package main

import (
	"github.com/joho/godotenv"
	"github.com/maevlava/ftf-clockify/internal/app"
	"github.com/maevlava/ftf-clockify/internal/config"
	httpdelivery "github.com/maevlava/ftf-clockify/internal/delivery/http"
	"github.com/maevlava/ftf-clockify/internal/service/workdebt"
	"log"
	"net/http"
)

const (
	port    = "3100"
	address = ":" + port
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.Load()
	appInstance := app.NewApp(cfg)

	// Services
	workDebtService := workdebt.NewService(cfg)

	// Handlers
	workDebtHandler := httpdelivery.NewWorkDebtHandler(workDebtService)

	router := httpdelivery.NewRouter(appInstance, workDebtHandler)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
