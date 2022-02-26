package main

import (
	"fmt"
	"foodlive/component"
	"foodlive/component/gosms"
	"foodlive/component/mycache"
	"foodlive/component/paymentprovider/momoprovider"
	"foodlive/component/tokenprovider/jwt"
	"foodlive/component/uploadprovider"
	"foodlive/config"
	"foodlive/pubsub/pubsublocal"
	"foodlive/subscribe"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func createDsnDb(username, password, host, port, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
}

func setupAppContext(appConfig *config.AppConfig) component.AppContext {
	//init database
	databaseDsn := createDsnDb(appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.DatabaseName)
	FDDatabase, err := gorm.Open(mysql.Open(databaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database notification- ", err)
	}
	FDDatabase = FDDatabase.Debug()

	//init cache
	myCache := mycache.NewMyCache()
	mySms := gosms.NewGoSms(appConfig.Sms.AccountSid, appConfig.Sms.AuthToken, appConfig.Sms.MyPhoneNumber)

	//init token provider
	tokenProvider := jwt.NewJwtProvider(appConfig.Token)

	//init upload provider
	s3Provider := uploadprovider.NewS3Provider(appConfig.S3AWS.BucketName, appConfig.S3AWS.Region, appConfig.S3AWS.ApiKey, appConfig.S3AWS.Secret, appConfig.S3AWS.Domain)

	//init pubsub local
	psLocal := pubsublocal.NewPubSubLocal()

	//init momo provider
	momoProvider := momoprovider.NewMomoProvider(appConfig.MomoConfig)

	//init app context
	appCtx := component.NewAppContext(appConfig, FDDatabase, myCache, mySms, tokenProvider, s3Provider, psLocal, momoProvider)

	//setup subscriber
	subscribe.SetupSubscriber(appCtx)

	return appCtx
}

func setupLog(appConfig *config.AppConfig) *os.File {
	f, err := os.OpenFile("food-hub.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("error opening file: %v", err)
	}
	log.SetOutput(f)
	//config log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if appConfig.Server.LevelLog >= 0 && appConfig.Server.LevelLog <= 6 {
		log.SetLevel(log.AllLevels[appConfig.Server.LevelLog])
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	return f
}
