package healthcheck

import (
	"github.com/cade-gray/vps-health-check/internal/logging"
)

func CheckAll(debug bool) bool {
	checkSuccess := false
	if debug {
		logging.InfoLogger.Println("Beginning to run All Health Checks")
	}
	apiSuccess := CheckAPI(debug)
	//CheckDB()
	//CheckServer()
	if apiSuccess /* && dbSuccess && serverSuccess */ {
		checkSuccess = true
		if debug {
			logging.InfoLogger.Println("All checks successful, checkSuccess has been set to true")
		}
	}
	if debug {
		logging.InfoLogger.Println("All Health Checks Completed. checkSuccess:", checkSuccess, " apiSuccess:", apiSuccess)
	}
	return checkSuccess
}
