package main

import (
	"flag"
	"log"
	"os"

	"github.com/cade-gray/vps-health-check/internal/alert"
	"github.com/cade-gray/vps-health-check/internal/config"
	"github.com/cade-gray/vps-health-check/internal/healthcheck"
	"github.com/cade-gray/vps-health-check/internal/logging"
	"github.com/joho/godotenv"
)

type FlagConfig struct {
	Debug bool
}

func main() {
	// Parse the command line flags
	debugPtr := flag.Bool("debug", false, "Debug mode")
	flag.Parse()
	flagConfig := FlagConfig{
		Debug: *debugPtr,
	}
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}
	// Load the configurations
	config.LoadConfig()
	// Initialize the logger
	logging.InitializeLogger("log.txt")
	// Log the start of the health check
	if flagConfig.Debug {
		logging.InfoLogger.Println("Health Check Starting")
	}
	// Perform the health check functions
	successTf := healthcheck.CheckAll(flagConfig.Debug)
	// Log the end of the health check if debug is enabled
	if flagConfig.Debug {
		logging.InfoLogger.Println("Health Check Ended")
	}
	// Log the success or failure of the health check
	if successTf {
		if flagConfig.Debug {
			logging.InfoLogger.Println("Health Check Successful")
			alert.Alert("Health Check Successful")
		}
		// Exit with a zero status code to indicate success
		os.Exit(0)
	} else {
		logging.ErrorLogger.Println("Health Check Failed")
		alert.Alert("Health Check FAILED!")
		// Exit with a non-zero status code to indicate failure
		os.Exit(1)
	}
}
