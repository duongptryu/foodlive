package restaurantownerstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
)

func (s *sqlStore) CreateOwnerRestaurant(ctx context.Context, data *restaurantownermodel.OwnerRestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
