package restaurantownerstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
)

func (s *sqlStore) UpdateStatusOwnerRestaurant(ctx context.Context, phoneNumber string) error {
	db := s.db

	if err := db.Table(restaurantownermodel.OwnerRestaurant{}.TableName()).Where("phone = ?", phoneNumber).Update("status", true).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) UpdateOwnerRestaurant(ctx context.Context, id int, data *restaurantownermodel.OwnerRestaurantUpdate) error {
	db := s.db

	if err := db.Table(restaurantownermodel.OwnerRestaurant{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
