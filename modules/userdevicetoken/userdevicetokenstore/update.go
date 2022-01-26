package userdevicetokenstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) UpdateUserDeviceToken(ctx context.Context, conditions map[string]interface{}, data *userdevicetokenmodel.UserDeviceTokenUpdate) error {
	db := s.db

	if err := db.Table(userdevicetokenmodel.UserDeviceTokenUpdate{}.TableName()).Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
