package useraddressbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/city/citymodel"
	"foodlive/modules/useraddress/useraddressmodel"
)

type createUserAddressBiz struct {
	cityStore        CityStore
	userAddressStore UserAddressStore
}

func NewCreateUserAddressBiz(cityStore CityStore, userAddressStore UserAddressStore) *createUserAddressBiz {
	return &createUserAddressBiz{
		cityStore:        cityStore,
		userAddressStore: userAddressStore,
	}
}

func (biz *createUserAddressBiz) CreateUserAddressBiz(ctx context.Context, data *useraddressmodel.UserAddressCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	cityDb, err := biz.cityStore.FindCity(ctx, map[string]interface{}{"id": data.CityId})
	if err != nil {
		return err
	}
	if cityDb.Id == 0 {
		return common.ErrDataNotFound(citymodel.EntityName)
	}

	if err := biz.userAddressStore.CreateUserAddress(ctx, data); err != nil {
		return err
	}

	return nil
}
