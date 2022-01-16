package userbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/gosms"
	"foodlive/component/mycache"
	"foodlive/modules/user/usermodel"
	log "github.com/sirupsen/logrus"
)

type ResendOTPActiveStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
}

type resendOTPActiveBiz struct {
	store   ResendOTPActiveStore
	myCache mycache.Cache
	mySms   gosms.GoSMS
}

func NewResendOTPActiveBiz(store ResendOTPActiveStore, myCache mycache.Cache, mySms gosms.GoSMS) *resendOTPActiveBiz {
	return &resendOTPActiveBiz{
		store:   store,
		myCache: myCache,
		mySms:   mySms,
	}
}

func (biz *resendOTPActiveBiz) ResendOTPActiveAccount(ctx context.Context, data *usermodel.ResendOTP) error {
	if err := data.Validate(); err != nil {
		return err
	}

	userDB, err := biz.store.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	if err != nil {
		return usermodel.ErUsernameOrPasswordInvalid
	}

	if userDB.Id == 0 {
		return usermodel.ErUsernameOrPasswordInvalid
	}

	if userDB.Status == true {
		return usermodel.ErrPhoneNumberAlreadyActivated
	}

	value, err := biz.myCache.Get(common.EntityOTP + data.Phone)
	if value != nil {
		return usermodel.ErrSendOTPMultiple
	}

	//Generate OTP
	OTP := common.GenerateOTP(4)

	//send OTP for user
	err = biz.mySms.SendOTP(ctx, data.Phone, OTP)
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	//set OTP in memory
	err = biz.myCache.SetWithExpire(common.EntityOTP+data.Phone, OTP, usermodel.TimeExpireOTPActivate)
	if err != nil {
		log.Error("Cannot set OTP to cache - ", err)
	}
	return nil
}
