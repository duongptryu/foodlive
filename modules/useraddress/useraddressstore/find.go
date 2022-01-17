package useraddressstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/useraddress/useraddressmodel"
)

func (s *sqlStore) FindUserAddressById(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*useraddressmodel.UserAddress, error) {
	db := s.db

	db = db.Table(useraddressmodel.UserAddress{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result useraddressmodel.UserAddress

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
