package useraddressbiz

import (
	"context"
	"foodlive/modules/useraddress/useraddressmodel"
)

type findUserAddress struct {
	userAddressStore UserAddressStore
}

func NewFindUserAddressBiz(userAddressStore UserAddressStore) *findUserAddress {
	return &findUserAddress{
		userAddressStore: userAddressStore,
	}
}

func (biz *findUserAddress) FindDefaultUserAddressBiz(ctx context.Context, userId int) (*useraddressmodel.UserAddress, error) {
	result, err := biz.userAddressStore.FindUserAddressById(ctx, map[string]interface{}{"user_id": userId, "status": true, "is_default": true})
	if err != nil {
		return nil, err
	}
	if result.Id == 0 {
		return nil, useraddressmodel.ErrUserDoseNotHaveDefaultAddress
	}

	return result, nil
}
