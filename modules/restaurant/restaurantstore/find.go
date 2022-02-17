package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result restaurantmodel.Restaurant

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
