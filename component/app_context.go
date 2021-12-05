package component

import (
	"fooddelivery/component/gosms"
	"fooddelivery/component/mycache"
	"fooddelivery/component/tokenprovider"
	"fooddelivery/config"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppConfig() *config.AppConfig
	GetDatabase() *gorm.DB
	GetMyCache() mycache.Cache
	GetMySms() gosms.GoSMS
	GetTokenProvider() tokenprovider.TokenProvider
}

type appCtx struct {
	appConfig     *config.AppConfig
	database      *gorm.DB
	myCache       mycache.Cache
	mySms         gosms.GoSMS
	tokenProvider tokenprovider.TokenProvider
}

func NewAppContext(appConfig *config.AppConfig, database *gorm.DB, myCache mycache.Cache, mySms gosms.GoSMS, tokenProvider tokenprovider.TokenProvider) *appCtx {
	return &appCtx{
		appConfig:     appConfig,
		database:      database,
		myCache:       myCache,
		mySms:         mySms,
		tokenProvider: tokenProvider,
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

func (ctx *appCtx) GetTokenProvider() tokenprovider.TokenProvider {
	return ctx.tokenProvider
}
