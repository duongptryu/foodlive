package useraddressbiz

import (
	"context"
	"foodlive/modules/city/citymodel"
	"foodlive/modules/useraddress/useraddressmodel"
)

type UserAddressStore interface {
	CreateUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error
	DeleteUserAddress(ctx context.Context, id int) error
	ListUserAddressByUserId(ctx context.Context, conditions map[string]interface{}, moreKey ...string) ([]useraddressmodel.UserAddress, error)
	UpdateUserAddress(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error
	FindUserAddressById(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*useraddressmodel.UserAddress, error)
}

type CityStore interface {
	FindCity(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*citymodel.City, error)
}
