package userdevicetokenstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) DeleteUserDeviceToken(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db

	if err := db.Table(userdevicetokenmodel.UserDeviceToken{}.TableName()).Where(conditions).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
