package rstcategorystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteRestaurantCategory(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
