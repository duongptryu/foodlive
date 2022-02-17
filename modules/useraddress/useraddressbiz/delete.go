package useraddressbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/useraddress/useraddressmodel"
)

type deleteUserAddressBiz struct {
	userAddressStore UserAddressStore
}

func NewDeleteUserAddressBiz(userAddressStore UserAddressStore) *deleteUserAddressBiz {
	return &deleteUserAddressBiz{
		userAddressStore: userAddressStore,
	}
}

func (biz *deleteUserAddressBiz) DeleteUserAddressBiz(ctx context.Context, id int) error {
	addrDb, err := biz.userAddressStore.FindUserAddressById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if addrDb.Id == 0 {
		return common.ErrDataNotFound(useraddressmodel.EntityName)
	}

	if err := biz.userAddressStore.DeleteUserAddress(ctx, id); err != nil {
		return err
	}

	return nil
}
