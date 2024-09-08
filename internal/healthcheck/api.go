package healthcheck

import (
	"io"
	"net/http"

	"github.com/cade-gray/vps-health-check/internal/config"
	"github.com/cade-gray/vps-health-check/internal/logging"
)

func CheckAPI() bool {
	// Implement API health check logic here
	logging.InfoLogger.Println("API health check Running")
	resp, err := http.Get(config.AppConfig.APIURL)
	if err != nil {
		logging.ErrorLogger.Println("Error checking API health:", err)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.ErrorLogger.Println("Error reading API health response:", err)
		return false
	}
	logging.InfoLogger.Println("API health check response:", string(body))
	if resp.StatusCode != http.StatusOK {
		logging.ErrorLogger.Println("API health check failed with status code:", resp.StatusCode)
		return false
	}

	logging.InfoLogger.Println("API health check Completed")
	return true
}
