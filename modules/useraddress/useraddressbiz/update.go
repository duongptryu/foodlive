package useraddressbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/city/citymodel"
	"foodlive/modules/useraddress/useraddressmodel"
)

type updateUserAddressBiz struct {
	cityStore        CityStore
	userAddressStore UserAddressStore
}

func NewUpdateUserAddressBiz(cityStore CityStore, userAddressStore UserAddressStore) *updateUserAddressBiz {
	return &updateUserAddressBiz{
		cityStore:        cityStore,
		userAddressStore: userAddressStore,
	}
}

func (biz *updateUserAddressBiz) UpdateUserAddressBiz(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	addrDb, err := biz.userAddressStore.FindUserAddressById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if addrDb.Id == 0 {
		return common.ErrDataNotFound(useraddressmodel.EntityName)
	}

	if addrDb.CityId != data.CityId && data.CityId != 0 {
		cityDb, err := biz.cityStore.FindCity(ctx, map[string]interface{}{"id": data.CityId})
		if err != nil {
			return err
		}
		if cityDb.Id == 0 {
			return common.ErrDataNotFound(citymodel.EntityName)
		}
	}

	if err := biz.userAddressStore.UpdateUserAddress(ctx, id, data); err != nil {
		return err
	}

	return nil
}
