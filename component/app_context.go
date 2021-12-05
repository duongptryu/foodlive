package component

import (
	"fooddelivery/component/gosms"
	"fooddelivery/component/mycache"
	"fooddelivery/config"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppConfig() *config.AppConfig
	GetDatabase() *gorm.DB
	GetMyCache() mycache.Cache
	GetMySms() gosms.GoSMS
}

type appCtx struct {
	appConfig        *config.AppConfig
	database   *gorm.DB
	myCache mycache.Cache
	mySms gosms.GoSMS
}

func NewAppContext (appConfig *config.AppConfig, database *gorm.DB, myCache mycache.Cache, mySms gosms.GoSMS) *appCtx {
	return &appCtx{
		appConfig: appConfig,
		database: database,
		myCache: myCache,
		mySms: mySms,
	}
}

func (ctx *appCtx) GetAppConfig() *config.AppConfig {
	return ctx.appConfig
}

func (ctx *appCtx) GetDatabase() *gorm.DB {
	return ctx.database
}

func (ctx *appCtx) GetMyCache() mycache.Cache {
	return ctx.myCache
}

func (ctx *appCtx) GetMySms() gosms.GoSMS {
	return ctx.mySms
}
