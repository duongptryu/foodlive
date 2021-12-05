package userbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component/gosms"
	"fooddelivery/component/mycache"
	"fooddelivery/modules/user/usermodel"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
}

type registerUserBiz struct {
	store   RegisterUserStore
	myCache mycache.Cache
	mySms   gosms.GoSMS
}

func NewRegisterUserBiz(store RegisterUserStore, myCache mycache.Cache, mySms gosms.GoSMS) *registerUserBiz {
	return &registerUserBiz{
		store:   store,
		myCache: myCache,
		mySms:   mySms,
	}
}

func (biz *registerUserBiz) RegisterUserBiz(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	userDB, err := biz.store.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	if userDB.Id != 0 {
		if userDB.Status == false {
			return usermodel.ErrPhoneNumberNotActivated
		}

		return usermodel.ErrPhoneNumberAlreadyExist
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	data.Password = string(hashedPassword)

	//create user in db
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	//Generate OTP
	OTP := usermodel.GenerateOTP(4)

	//send OTP for user
	err = biz.mySms.SendOTP(ctx, data.Phone, OTP)
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	//set OTP in memory
	err = biz.myCache.Set(usermodel.EntityOTP+OTP, data.Phone)
	if err != nil {
		log.Error("Cannot set OTP to cache - ", err)
	}

	return nil
}
