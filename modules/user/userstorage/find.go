package userstorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/user/usermodel"
)

func (s *SQLStore) FindUser (ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error) {
	db := s.db
	var result usermodel.User

	db = db.Where(conditions)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}