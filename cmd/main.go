package main

import (
	"github.com/joho/godotenv"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/http/router"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(".env file couldn't be loaded")
	}

	// Init Router
	router.Initialize()
}
