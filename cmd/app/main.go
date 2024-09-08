package main

import (
	"fmt"

	"github.com/cade-gray/vps-health-check/internal/healthcheck"
	"github.com/cade-gray/vps-health-check/internal/logging"
)

func main() {
	fmt.Println("Hello, World!")
	logging.InitializeLogger("log.txt")
	logging.Logger.Println("Health Check Starting")
	healthcheck.CheckAPI()
	logging.Logger.Println("Health Check Ended")
}
