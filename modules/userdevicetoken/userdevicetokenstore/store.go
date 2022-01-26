package userdevicetokenstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
	"gorm.io/gorm"
)

type UserDeviceTokenStore interface {
	CreateUserDeviceToken(ctx context.Context, data *userdevicetokenmodel.UserDeviceTokenCreate) error
	DeleteUserDeviceToken(ctx context.Context, conditions map[string]interface{}) error
	FindUserDeviceToken(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*userdevicetokenmodel.UserDeviceToken, error)
	ListUserDeviceToken(ctx context.Context,
		condition map[string]interface{},
		filter *userdevicetokenmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]userdevicetokenmodel.UserDeviceToken, error)
	UpdateUserDeviceToken(ctx context.Context, condition map[string]interface{}, data *userdevicetokenmodel.UserDeviceTokenUpdate) error
}

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}
