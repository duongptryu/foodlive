package userdevicetokenstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) FindUserDeviceToken(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*userdevicetokenmodel.UserDeviceToken, error) {
	db := s.db
	var result userdevicetokenmodel.UserDeviceToken

	db = db.Where(conditions)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
