package restaurantratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) FindRestaurantRating(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantratingmodel.RestaurantRating, error) {
	db := s.db

	db = db.Table(restaurantratingmodel.RestaurantRating{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result restaurantratingmodel.RestaurantRating

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
