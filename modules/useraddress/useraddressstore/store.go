package useraddressstore

import (
	"context"
	"foodlive/modules/useraddress/useraddressmodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type UserAddressStore interface {
	CreateUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error
	DeleteUserAddress(ctx context.Context, id int) error
	ListUserAddressByUserId(ctx context.Context, conditions map[string]interface{}, moreKey ...string) ([]useraddressmodel.UserAddress, error)
	UpdateUserAddress(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error
	FindUserAddressById(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*useraddressmodel.UserAddress, error)
}
