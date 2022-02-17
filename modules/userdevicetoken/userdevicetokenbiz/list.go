package userdevicetokenbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
)

type listUserDeviceTokenBiz struct {
	store userdevicetokenstore.UserDeviceTokenStore
}

func NewListUserDeviceTokenBiz(store userdevicetokenstore.UserDeviceTokenStore) *listUserDeviceTokenBiz {
	return &listUserDeviceTokenBiz{store: store}
}

func (biz *listUserDeviceTokenBiz) ListUserDeviceTokenBiz(ctx context.Context, paging *common.Paging, filter *userdevicetokenmodel.Filter) ([]userdevicetokenmodel.UserDeviceToken, error) {
	result, err := biz.store.ListUserDeviceToken(ctx, nil, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
