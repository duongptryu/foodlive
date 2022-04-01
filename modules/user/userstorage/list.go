package userstorage

import (
	"context"
	"foodlive/common"
	"foodlive/modules/user/usermodel"
)

func (s *sqlStore) ListUserWithoutPaging(ctx context.Context,
	condition map[string]interface{},
	filter *usermodel.Filter,
	moreKey ...string,
) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db

	db = db.Table(usermodel.User{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.CreatedAtGt != nil {
			db = db.Where("created_at > ?", *v.CreatedAtGt).Order("created_at ASC")
		}
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
