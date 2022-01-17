package useraddressbiz

import (
	"context"
	"foodlive/modules/useraddress/useraddressmodel"
)

type listUserAddressBiz struct {
	userAddressStore UserAddressStore
}

func NewListUserAddressBiz(userAddressStore UserAddressStore) *listUserAddressBiz {
	return &listUserAddressBiz{
		userAddressStore: userAddressStore,
	}
}

func (biz *listUserAddressBiz) ListUserAddressBiz(ctx context.Context, userId int) ([]useraddressmodel.UserAddress, error) {
	result, err := biz.userAddressStore.ListUserAddressByUserId(ctx, map[string]interface{}{"user_id": userId, "status": true})
	if err != nil {
		return nil, err
	}

	return result, nil
}
