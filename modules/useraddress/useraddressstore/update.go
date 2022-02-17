package useraddressstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/useraddress/useraddressmodel"
)

func (s *sqlStore) UpdateUserAddress(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error {
	db := s.db

	if err := db.Table(useraddressmodel.UserAddressUpdate{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
