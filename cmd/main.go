package main

import (
	"github.com/joho/godotenv"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/http/router"
)

func main() {
	// Load .env file
	godotenv.Load()

	// Init Router
	router.Initialize()
}
