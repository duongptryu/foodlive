package userbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component/mycache"
	"fooddelivery/modules/user/usermodel"
	log "github.com/sirupsen/logrus"
)

type ActivateUserStore interface {
	UpdateStatusUser(ctx context.Context, phoneNumber string) error
}

type activateUserBiz struct {
	store   ActivateUserStore
	myCache mycache.Cache
}

func NewActiveUserBiz(store ActivateUserStore, myCache mycache.Cache) *activateUserBiz {
	return &activateUserBiz{
		store:   store,
		myCache: myCache,
	}
}

func (biz *activateUserBiz) UserActiveBiz(ctx context.Context, data *usermodel.UserActive) error {
	phoneNumber, err := biz.myCache.Get(usermodel.EntityOTP + data.OTP)
	if err != nil {
		return usermodel.ErrOTPInvalidOrExpire
	}

	if phoneNumber.(string) != data.Phone {
		return usermodel.ErrOTPInvalidOrExpire
	}

	//remove cache
	if err := biz.myCache.Remove(usermodel.EntityOTP + data.OTP); err != nil {
		log.Error("Cannot remove OTP in cache - ", err)
	}

	if err := biz.store.UpdateStatusUser(ctx, data.Phone); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
