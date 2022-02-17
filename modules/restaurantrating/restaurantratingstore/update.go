package restaurantratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) UpdateRestaurantRating(ctx context.Context, id int, data *restaurantratingmodel.RestaurantRatingUpdate) error {
	db := s.db

	if err := db.Table(restaurantratingmodel.RestaurantRating{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
