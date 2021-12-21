package foodstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/food/foodmodel"
)

func (s *sqlStore) FindFood(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result foodmodel.Food

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
