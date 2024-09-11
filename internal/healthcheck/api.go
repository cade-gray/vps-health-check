package healthcheck

import (
	"io"
	"net/http"

	"github.com/cade-gray/vps-health-check/internal/config"
	"github.com/cade-gray/vps-health-check/internal/logging"
)

func CheckAPI(debug bool) bool {
	// Log the start of the API health check if debug is enabled
	if debug {
		logging.InfoLogger.Println("API health check Running")
		logging.InfoLogger.Println("Attempting GET to API URL: ", config.AppConfig.APIURL)
	}
	// Perform the GET request to the API
	resp, err := http.Get(config.AppConfig.APIURL)
	// Check if there was an error performing the GET request
	if err != nil {
		logging.ErrorLogger.Println("Error checking API health: ", err)
		return false
	}
	// Defer closing the response body until the function completes
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	// Check if there was an error reading the response body and return false if there was
	if err != nil {
		logging.ErrorLogger.Println("Error reading API health response: ", err)
		return false
	}
	// Check if the status code is not 200 and return false if it is not
	if resp.StatusCode != http.StatusOK {
		logging.ErrorLogger.Println("API health check failed with status code: ", resp.StatusCode)
		return false
	}
	// Log the response body and health check completion if debug is enabled
	if debug {
		logging.InfoLogger.Println("API health check response body:", string(body))
		logging.InfoLogger.Println("API health check Completed")
	}
	return true
}
