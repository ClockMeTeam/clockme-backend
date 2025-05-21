package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maevlava/ftf-clockify/internal/app"
	"github.com/maevlava/ftf-clockify/internal/config"
	httpdelivery "github.com/maevlava/ftf-clockify/internal/delivery/http"
	"github.com/maevlava/ftf-clockify/internal/service/workdebt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
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
	go openBrowser(fmt.Sprintf("http://localhost%s/api/v1/debts", address))
	log.Fatal(server.ListenAndServe())
}
func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		err := exec.Command("xdg-open", url).Start()
		if err != nil {
			return err
		}
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}

	return nil
}
