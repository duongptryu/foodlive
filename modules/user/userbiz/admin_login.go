package userbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"foodlive/modules/user/usermodel"
)

type adminLoginBiz struct {
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewAdminLogin(tokenProvider tokenprovider.TokenProvider, expiry int) *loginBusiness {
	return &loginBusiness{
		tokenProvider: tokenProvider,
		expire:        expiry,
	}
}

func (biz *loginBusiness) AdminLoginBiz(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	if data.Phone != "admin" {
		return nil, common.ErrUnAuthorization
	}
	if data.Password != "123123" {
		return nil, common.ErrUnAuthorization
	}

	payload := tokenprovider.TokenPayload{
		UserId: 0,
		Role:   "admin",
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, nil)

	return account, nil
}
