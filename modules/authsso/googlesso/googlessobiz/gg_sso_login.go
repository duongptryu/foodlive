package googlessobiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"foodlive/modules/authsso/googlesso/googlessomodel"
)

type loginGoogleBiz struct {
	store         GoogleStore
	userStore     UserStore
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewLoginGoogleBiz(store GoogleStore, userStore UserStore, tokenProvider tokenprovider.TokenProvider, expire int) *loginGoogleBiz {
	return &loginGoogleBiz{
		store:         store,
		userStore:     userStore,
		tokenProvider: tokenProvider,
		expire:        expire,
	}
}

func (biz *loginGoogleBiz) LoginGoogleBiz(ctx context.Context, data *googlessomodel.GoogleJwtInput) (*googlessomodel.Account, error) {
	result, err := biz.store.ValidateGoogleJwt(ctx, data.Token)
	if err != nil {
		return nil, err
	}

	userDB, err := biz.userStore.FindUser(ctx, map[string]interface{}{"gg_id": result.Subject})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	if userDB.Id == 0 {
		return nil, googlessomodel.ErrAccountDoesNotExist
	}

	payload := tokenprovider.TokenPayload{
		UserId: userDB.Id,
		Role:   userDB.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := googlessomodel.NewAccount(accessToken, nil)

	return account, nil
}
