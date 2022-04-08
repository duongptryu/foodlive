package main

import (
	"foodlive/config"
	"foodlive/eventsmartcontract"
	pb "foodlive/gen/proto"
	"foodlive/modules/hello/hellotransport/grpchello"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	eventWatcher := eventsmartcontract.NewEventWatcher(appConfig)
	go eventWatcher.Watch(appCtx)

	r := gin.Default()
	r.Use(cors.Default())

	setupRouter(r, appCtx)

	//run grpc server
	go func() {
		mux := http.NewServeMux()
		helloServerImpl := grpchello.ServerHello{}
		helloHandle := pb.NewHelloServer(&helloServerImpl)
		mux.Handle(helloHandle.PathPrefix(), helloHandle)

		http.ListenAndServe(":8081", mux)
	}()

	r.Run(":" + appConfig.Server.Port)
}
