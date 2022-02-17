package userbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/mycache"
	"foodlive/modules/user/usermodel"
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
	otp, err := biz.myCache.Get(common.EntityOTP + data.Phone)
	if err != nil {
		return usermodel.ErrOTPInvalidOrExpire
	}

	if otp.(string) != data.OTP {
		return usermodel.ErrOTPInvalidOrExpire
	}

	//remove cache
	if err := biz.myCache.Remove(common.EntityOTP + data.Phone); err != nil {
		log.Error("Cannot remove OTP in cache - ", err)
	}

	if err := biz.store.UpdateStatusUser(ctx, data.Phone); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
