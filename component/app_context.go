package component

import (
	"fooddelivery/config"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppConfig() *config.AppConfig
	GetDatabase() *gorm.DB
}

type appCtx struct {
	appConfig        *config.AppConfig
	database   *gorm.DB
}

func NewAppContext (appConfig *config.AppConfig, database *gorm.DB) *appCtx {
	return &appCtx{
		appConfig: appConfig,
		database: database,
	}
}

func (ctx *appCtx) GetAppConfig() *config.AppConfig {
	return ctx.appConfig
}

func (ctx *appCtx) GetDatabase() *gorm.DB {
	return ctx.database
}
