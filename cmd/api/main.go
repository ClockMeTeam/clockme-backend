package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maevlava/ftf-clockify/internal/app"
	"github.com/maevlava/ftf-clockify/internal/config"
	httpdelivery "github.com/maevlava/ftf-clockify/internal/delivery/http"
	"log"
	"net/http"
)

const (
	port    = "8080"
	address = ":" + port
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.Load()
	app := app.NewApp(cfg)

	router := httpdelivery.NewRouter(app)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	fmt.Println("Starting server on port ", port)
	log.Fatal(server.ListenAndServe())
}
