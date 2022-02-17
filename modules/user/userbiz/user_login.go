package userbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"foodlive/modules/user/usermodel"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type LoginHashProvider interface {
	ValidatePassword(password, hasedPassword string) (bool, error)
}

type loginBusiness struct {
	storeUser     UserLoginStore
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewLoginBiz(storeUser UserLoginStore, tokenProvider tokenprovider.TokenProvider, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		expire:        expiry,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	userDB, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	if err != nil {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}

	if userDB.Id == 0 {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}

	if userDB.Status == false {
		return nil, usermodel.ErrPhoneNumberNotActivated
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(data.Password))
	if err != nil {
		return nil, usermodel.ErUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: userDB.Id,
		Role:   userDB.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)

	if err != nil {
		return nil, common.ErrInternal(err)
	}
	//
	//refreshToken, err := biz.tokenProvider.Generate(&payload, biz.expire)
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	account := usermodel.NewAccount(accessToken, nil)

	return account, nil
}
