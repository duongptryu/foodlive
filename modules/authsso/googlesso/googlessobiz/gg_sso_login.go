package googlessobiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"foodlive/modules/authsso/googlesso/googlessomodel"
	"foodlive/modules/user/usermodel"
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
	var payload tokenprovider.TokenPayload
	if userDB.Id == 0 {
		// gg_id dose not exist
		userCreate := usermodel.UserCreate{
			Status:    false,
			Role:      "user",
			Phone:     "",
			LastName:  result.LastName,
			FirstName: result.FirstName,
			Password:  "",
			GgId:      &result.Subject,
		}
		if err := biz.userStore.CreateUser(ctx, &userCreate); err != nil {
			return nil, common.ErrCannotCreateEntity(googlessomodel.EntityName, err)
		}
		payload = tokenprovider.TokenPayload{
			UserId: userCreate.Id,
			Role:   userCreate.Role,
			Type:   common.TypeAccountSocial,
		}
	} else {
		payload = tokenprovider.TokenPayload{
			UserId: userDB.Id,
			Role:   userDB.Role,
			Type:   common.TypeAccountSocial,
		}
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := googlessomodel.NewAccount(accessToken, nil)

	return account, nil
}
