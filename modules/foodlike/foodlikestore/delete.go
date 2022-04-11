package foodlikestore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
)

func (s *sqlStore) DeleteFoodLike(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(foodlikemodel.FoodLike{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
