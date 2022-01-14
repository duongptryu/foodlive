package categorystore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/category/categorymodel"
)

func (s *sqlStore) FindCategory(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*categorymodel.Category, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result categorymodel.Category

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
