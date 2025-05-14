package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//cfg := config.Load()

}
