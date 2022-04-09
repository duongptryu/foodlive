package foodratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodrating/foodratingmodel"
)

func (s *sqlStore) FindFoodRating(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*foodratingmodel.FoodRating, error) {
	db := s.db

	db = db.Table(foodratingmodel.FoodRating{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result foodratingmodel.FoodRating

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
