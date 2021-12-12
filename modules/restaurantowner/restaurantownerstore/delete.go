package restaurantownerstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
)

func (s *sqlStore) DeleteOwnerRestaurant(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantownermodel.OwnerRestaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
