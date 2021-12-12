package restaurantownerstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
)

func (s *sqlStore) UpdateStatusOwnerRestaurant(ctx context.Context, phoneNumber string) error {
	db := s.db

	if err := db.Table(restaurantownermodel.OwnerRestaurant{}.TableName()).Where("phone = ?", phoneNumber).Update("status", true).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
