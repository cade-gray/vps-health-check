package healthcheck

func CheckAll() bool {
	checkSuccess := false
	apiSuccess := CheckAPI()
	//CheckDB()
	//CheckServer()
	if apiSuccess /* && dbSuccess && serverSuccess */ {
		checkSuccess = true
	}
	return checkSuccess
}
