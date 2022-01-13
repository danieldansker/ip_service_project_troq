package main

import (
	"os"

	"github.com/danieldansker/ip_service_project_troq/app"
	"github.com/danieldansker/ip_service_project_troq/utils"
)

func main() {
	os.Setenv("DB_URL", "DBReaderMockData")
	os.Setenv("RATE_LIMITER", "2")
	configuration := utils.GetConfiguration()
	app := &app.App{}
	app.InitApp(configuration)
	app.RunApp()
}
