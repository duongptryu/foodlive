package userstorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
