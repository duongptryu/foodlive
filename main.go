package main

import (
	"foodlive/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	runService()
}

func runService() {
	//init config
	appConfig, err := config.NewAppConfig("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	//config mode
	if appConfig.Server.ModeRun == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	f := setupLog(appConfig)
	defer f.Close()

	appCtx := setupAppContext(appConfig)

	r := gin.Default()
	r.Use(cors.Default())

	setupRouter(r, appCtx)

	r.Run(":" + appConfig.Server.Port)

}
