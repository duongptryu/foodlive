package restaurantratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) CreateRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
