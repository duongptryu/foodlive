package component

import (
	"foodlive/component/gosms"
	"foodlive/component/mycache"
	"foodlive/component/notification"
	"foodlive/component/paymentprovider"
	"foodlive/component/tokenprovider"
	"foodlive/component/uploadprovider"
	"foodlive/config"
	"foodlive/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppConfig() *config.AppConfig
	GetDatabase() *gorm.DB
	GetMyCache() mycache.Cache
	GetMySms() gosms.GoSMS
	GetTokenProvider() tokenprovider.TokenProvider
	GetUploadProvider() uploadprovider.UploadProvider
	GetPubSubProvider() pubsub.PubSub
	GetPaymentProvider() paymentprovider.PaymentProvider
	GetCryptoPaymentProvider() paymentprovider.CryptoPaymentProvider
}

type appCtx struct {
	appConfig             *config.AppConfig
	database              *gorm.DB
	myCache               mycache.Cache
	mySms                 gosms.GoSMS
	tokenProvider         tokenprovider.TokenProvider
	uploadProvider        uploadprovider.UploadProvider
	pubSubProvider        pubsub.PubSub
	paymentProvider       paymentprovider.PaymentProvider
	cryptoPaymentProvider paymentprovider.CryptoPaymentProvider
	notiProvider          notification.NotiProvider
}

func NewAppContext(appConfig *config.AppConfig, database *gorm.DB, myCache mycache.Cache, mySms gosms.GoSMS, tokenProvider tokenprovider.TokenProvider, uploadProvider uploadprovider.UploadProvider, pubSubProvider pubsub.PubSub, paymentProvider paymentprovider.PaymentProvider, cryptoPaymentProvider paymentprovider.CryptoPaymentProvider, notiProvider notification.NotiProvider) *appCtx {
	return &appCtx{
		appConfig:             appConfig,
		database:              database,
		myCache:               myCache,
		mySms:                 mySms,
		tokenProvider:         tokenProvider,
		uploadProvider:        uploadProvider,
		pubSubProvider:        pubSubProvider,
		paymentProvider:       paymentProvider,
		cryptoPaymentProvider: cryptoPaymentProvider,
		notiProvider:          notiProvider,
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

func (ctx *appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetPubSubProvider() pubsub.PubSub {
	return ctx.pubSubProvider
}

func (ctx *appCtx) GetPaymentProvider() paymentprovider.PaymentProvider {
	return ctx.paymentProvider
}

func (ctx *appCtx) GetCryptoPaymentProvider() paymentprovider.CryptoPaymentProvider {
	return ctx.cryptoPaymentProvider
}

func (ctx *appCtx) GetNotiProvider() notification.NotiProvider {
	return ctx.notiProvider
}
