package useraddressstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/useraddress/useraddressmodel"
)

func (s *sqlStore) ListUserAddressByUserId(ctx context.Context, conditions map[string]interface{}, moreKey ...string) ([]useraddressmodel.UserAddress, error) {
	db := s.db
	var result []useraddressmodel.UserAddress

	db = db.Where(conditions)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
