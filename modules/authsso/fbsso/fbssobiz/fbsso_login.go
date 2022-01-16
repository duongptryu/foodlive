package fbssobiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"foodlive/modules/authsso/fbsso/fbssomodel"
)

type loginFbBiz struct {
	store         FbStore
	userStore     UserStore
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewLoginFbBiz(store FbStore, userStore UserStore, tokenProvider tokenprovider.TokenProvider, expire int) *loginFbBiz {
	return &loginFbBiz{
		store:         store,
		userStore:     userStore,
		tokenProvider: tokenProvider,
		expire:        expire,
	}
}

func (biz *loginFbBiz) LoginFbBiz(ctx context.Context, data *fbssomodel.FacebookJwtInput) (*fbssomodel.Account, error) {
	result, err := biz.store.ValidateFbJwt(ctx, data.Token)
	if err != nil {
		return nil, err
	}

	userDB, err := biz.userStore.FindUser(ctx, map[string]interface{}{"fb_id": result.ID})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	if userDB.Id == 0 {
		return nil, fbssomodel.ErrAccountDoesNotExist
	}

	payload := tokenprovider.TokenPayload{
		UserId: userDB.Id,
		Role:   userDB.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := fbssomodel.NewAccount(accessToken, nil)

	return account, nil
}
