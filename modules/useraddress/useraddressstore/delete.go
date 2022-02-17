package useraddressstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/useraddress/useraddressmodel"
)

func (s *sqlStore) DeleteUserAddress(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(useraddressmodel.UserAddress{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
