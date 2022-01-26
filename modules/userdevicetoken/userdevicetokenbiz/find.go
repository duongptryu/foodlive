package userdevicetokenbiz

import (
	"context"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
)

type findUserDeviceTokenBiz struct {
	store userdevicetokenstore.UserDeviceTokenStore
}

func NewFindUserDeviceTokenBiz(store userdevicetokenstore.UserDeviceTokenStore) *findUserDeviceTokenBiz {
	return &findUserDeviceTokenBiz{store: store}
}

func (biz *findUserDeviceTokenBiz) FindUserDeviceTokenBiz(ctx context.Context, userId int) (*userdevicetokenmodel.UserDeviceToken, error) {
	result, err := biz.store.FindUserDeviceToken(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	return result, nil
}
