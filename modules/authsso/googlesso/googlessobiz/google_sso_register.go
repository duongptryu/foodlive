package googlessobiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/authsso/googlesso/googlessomodel"
	"fooddelivery/modules/user/usermodel"
)

type registerGoogleBiz struct {
	store     GoogleStore
	userStore UserStore
}

func NewRegisterGoogleBiz(store GoogleStore, userStore UserStore) *registerGoogleBiz {
	return &registerGoogleBiz{
		store:     store,
		userStore: userStore,
	}
}

func (biz *registerGoogleBiz) RegisterGoogleBiz(ctx context.Context, data *googlessomodel.GoogleJwtInput) (string, error) {
	result, err := biz.store.ValidateGoogleJwt(ctx, data.Token)
	if err != nil {
		return "", err
	}

	userDB, err := biz.userStore.FindUser(ctx, map[string]interface{}{"gg_id": result.Subject})
	if err != nil {
		return "", common.ErrDB(err)
	}
	if userDB.Id != 0 {
		return "", googlessomodel.ErrAccountAlreadyExisted
	}

	userCreate := usermodel.UserCreate{
		Status:    false,
		Role:      "user",
		Phone:     "",
		LastName:  result.LastName,
		FirstName: result.FirstName,
		Password:  "",
		GgId:      result.Subject,
	}

	if err := biz.userStore.CreateUser(ctx, &userCreate); err != nil {
		return "", common.ErrCannotCreateEntity(googlessomodel.EntityName, err)
	}

	return "", nil
}
