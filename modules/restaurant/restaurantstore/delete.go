package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteRestaurant(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
