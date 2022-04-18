package userbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/gosms"
	"foodlive/component/mycache"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/user/userstorage"
	log "github.com/sirupsen/logrus"
)

type updateUserBiz struct {
	store userstorage.UserStore
}

func NewUpdateUserBiz(store userstorage.UserStore) *updateUserBiz {
	return &updateUserBiz{
		store: store,
	}
}

func (biz *updateUserBiz) AdminUpdateUserBiz(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	if err := biz.store.UpdateUser(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}

func (biz *updateUserBiz) UpdateMyProfule(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	if err := biz.store.UpdateUser(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}

func (biz *updateUserBiz) UpdateUserBiz(ctx context.Context, id int, data *usermodel.UserUpdate, myCache mycache.Cache, mySms gosms.GoSMS) error {
	if err := data.Validate(); err != nil {
		return err
	}
	//check if phone number exist in db
	exist, err := biz.store.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	if err != nil {
		return err
	}
	if exist.Id != 0 {
		return usermodel.ErrPhoneNumberAlreadyExist
	}

	if err := biz.store.UpdateUser(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	//Generate OTP
	OTP := common.GenerateOTP(4)

	//send OTP for user
	err = mySms.SendOTP(ctx, data.Phone, OTP)
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	//set OTP in memory
	err = myCache.SetWithExpire(common.EntityOTP+data.Phone, OTP, usermodel.TimeExpireOTPActivate)
	if err != nil {
		log.Error("Cannot set OTP to cache - ", err)
	}

	return nil
}
