package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
