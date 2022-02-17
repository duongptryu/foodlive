package useraddressstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/useraddress/useraddressmodel"
)

func (s *sqlStore) CreateUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
