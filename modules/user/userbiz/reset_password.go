package userbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component/mycache"
	"fooddelivery/modules/user/usermodel"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserResetPassword interface {
	UpdatePasswordUser(ctx context.Context, data *usermodel.UserResetPassword) error
}

type userResetPasswordBiz struct {
	store   UserResetPassword
	myCache mycache.Cache
}

func NewUserResetPasswordBiz(store UserResetPassword, myCache mycache.Cache) *userResetPasswordBiz {
	return &userResetPasswordBiz{
		store:   store,
		myCache: myCache,
	}
}

func (biz *userResetPasswordBiz) UserResetPasswordBiz(ctx context.Context, data *usermodel.UserResetPassword) error {
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

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	data.Password = string(hashedPassword)

	if err := biz.store.UpdatePasswordUser(ctx, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
