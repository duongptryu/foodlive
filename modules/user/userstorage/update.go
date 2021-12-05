package userstorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/user/usermodel"
)

func (s *SQLStore) UpdateStatusUser(ctx context.Context, phoneNumber string) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("phone = ?", phoneNumber).Update("status", true).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
