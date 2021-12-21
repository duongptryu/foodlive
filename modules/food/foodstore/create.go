package foodstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/food/foodmodel"
)

func (s *sqlStore) CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
