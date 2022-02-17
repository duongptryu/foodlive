package fbssobiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/authsso/fbsso/fbssomodel"
	"foodlive/modules/user/usermodel"
)

type registerFbSsoStore struct {
	store     FbStore
	userStore UserStore
}

func NewRegisterFbSsoBiz(store FbStore, userStore UserStore) *registerFbSsoStore {
	return &registerFbSsoStore{
		store:     store,
		userStore: userStore,
	}
}

func (biz *registerFbSsoStore) RegisterFbSsoBiz(ctx context.Context, data *fbssomodel.FacebookJwtInput) (string, error) {
	result, err := biz.store.ValidateFbJwt(ctx, data.Token)
	if err != nil {
		return "", err
	}

	userDB, err := biz.userStore.FindUser(ctx, map[string]interface{}{"fb_id": result.ID})
	if err != nil {
		return "", common.ErrDB(err)
	}
	if userDB.Id != 0 {
		return "", fbssomodel.ErrAccountAlreadyExisted
	}

	userCreate := usermodel.UserCreate{
		Status:    false,
		Role:      "user",
		Phone:     "",
		LastName:  result.Name,
		FirstName: result.Name,
		Password:  "",
		FbId:      result.ID,
	}

	if err := biz.userStore.CreateUser(ctx, &userCreate); err != nil {
		return "", common.ErrCannotCreateEntity(fbssomodel.EntityName, err)
	}

	return "", nil
}
