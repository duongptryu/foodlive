package userstorage

import (
	"context"
	"foodlive/common"
	"foodlive/modules/user/usermodel"
)

func (s *sqlStore) UpdateStatusUser(ctx context.Context, phoneNumber string) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("phone = ?", phoneNumber).Update("status", true).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) UpdatePasswordUser(ctx context.Context, data *usermodel.UserResetPassword) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("phone = ?", data.Phone).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
