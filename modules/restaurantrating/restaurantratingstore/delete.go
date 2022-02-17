package restaurantratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) DeleteRestaurantRating(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantratingmodel.RestaurantRating{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
