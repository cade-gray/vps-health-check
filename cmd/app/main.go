package main

import (
	"log"
	"os"

	"github.com/cade-gray/vps-health-check/internal/alert"
	"github.com/cade-gray/vps-health-check/internal/config"
	"github.com/cade-gray/vps-health-check/internal/healthcheck"
	"github.com/cade-gray/vps-health-check/internal/logging"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		logging.ErrorLogger.Println("No .env file found")
		return
	}
	// Load the configuration
	config.LoadConfig()
	// Initialize the logger
	logging.InitializeLogger("log.txt")
	// Log the start of the health check
	logging.InfoLogger.Println("Health Check Starting")
	// Perform the health check functions
	successTf := healthcheck.CheckAll()
	// Log the end of the health check
	logging.InfoLogger.Println("Health Check Ended")
	// Log the success or failure of the health check
	if successTf {
		logging.InfoLogger.Println("Health Check Successful")
		alert.Alert("Health Check Successful")
	} else {
		logging.ErrorLogger.Println("Health Check Failed")
		alert.Alert("Health Check FAILED!")
		// Exit with a non-zero status code to indicate failure
		os.Exit(1)
	}
}
