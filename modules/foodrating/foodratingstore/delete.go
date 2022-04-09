package foodratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodrating/foodratingmodel"
)


func (s *sqlStore) DeleteFoodRating(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(foodratingmodel.FoodRating{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
