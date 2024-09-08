package main

import (
	"log"

	"github.com/cade-gray/vps-health-check/internal/config"
	"github.com/cade-gray/vps-health-check/internal/healthcheck"
	"github.com/cade-gray/vps-health-check/internal/logging"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Load the configuration
	config.LoadConfig()
	// Initialize the logger
	logging.InitializeLogger("log.txt")
	// Log the start of the health check
	logging.Logger.Println("Health Check Starting")
	// Perform the health check functions
	healthcheck.CheckAPI()
	// Log the end of the health check
	logging.Logger.Println("Health Check Ended")
}
