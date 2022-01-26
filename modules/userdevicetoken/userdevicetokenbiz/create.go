package userdevicetokenbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
)

type createUserDeviceTokenBiz struct {
	store userdevicetokenstore.UserDeviceTokenStore
}

func NewCreateUserDeviceTokenBiz(store userdevicetokenstore.UserDeviceTokenStore) *createUserDeviceTokenBiz {
	return &createUserDeviceTokenBiz{
		store: store,
	}
}

func (biz *createUserDeviceTokenBiz) CreateUserDeviceTokenBiz(ctx context.Context, data *userdevicetokenmodel.UserDeviceTokenCreate) error {
	exist, err := biz.store.FindUserDeviceToken(ctx, map[string]interface{}{"user_id": data.UserId})
	if err != nil {
		return err
	}

	if exist.Id == 0 {
		// doesn't exist
		if err := biz.store.CreateUserDeviceToken(ctx, data); err != nil {
			return err
		}

		return nil
	}
	//exist -> update
	dataUpdate := userdevicetokenmodel.UserDeviceTokenUpdate{
		Token: data.Token,
		Os:    data.Os,
	}
	if err := biz.store.UpdateUserDeviceToken(ctx, map[string]interface{}{"id": exist.Id}, &dataUpdate); err != nil {
		return common.ErrCannotCreateEntity(userdevicetokenmodel.EntityName, err)
	}

	return nil
}
