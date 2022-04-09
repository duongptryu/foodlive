package foodratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodrating/foodratingmodel"
)


func (s *sqlStore) CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
