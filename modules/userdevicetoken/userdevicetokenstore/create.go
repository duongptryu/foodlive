package userdevicetokenstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) CreateUserDeviceToken(ctx context.Context, data *userdevicetokenmodel.UserDeviceTokenCreate) error {
	db := s.db

	data.Status = true

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
