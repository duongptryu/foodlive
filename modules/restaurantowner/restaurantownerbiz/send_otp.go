package restaurantownerbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component/gosms"
	"fooddelivery/component/mycache"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
	"fooddelivery/modules/user/usermodel"
	log "github.com/sirupsen/logrus"
)

type SendOTPActiveStore interface {
	FindOwnerRestaurant(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*restaurantownermodel.OwnerRestaurant, error)
}

type sendOTPActiveBiz struct {
	store   SendOTPActiveStore
	myCache mycache.Cache
	mySms   gosms.GoSMS
}

func NewSendOTPActiveBiz(store SendOTPActiveStore, myCache mycache.Cache, mySms gosms.GoSMS) *sendOTPActiveBiz {
	return &sendOTPActiveBiz{
		store:   store,
		myCache: myCache,
		mySms:   mySms,
	}
}

func (biz *sendOTPActiveBiz) SendOTPActiveBiz(ctx context.Context, data *restaurantownermodel.SendOTP) error {
	if err := data.Validate(); err != nil {
		return err
	}

	userDB, err := biz.store.FindOwnerRestaurant(ctx, map[string]interface{}{"phone": data.Phone})
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
