package useraddressstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/useraddress/useraddressmodel"
)

func (s *sqlStore) UpdateUserAddress(ctx context.Context, id int, data *useraddressmodel.UserAddressUpdate) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
