package foodratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodrating/foodratingmodel"
)

func (s *sqlStore) UpdateFoodRating(ctx context.Context, id int, data *foodratingmodel.FoodRatingUpdate) error {
	db := s.db

	if err := db.Table(foodratingmodel.FoodRating{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}