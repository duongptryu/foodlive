package restaurantstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db

	db = db.Where(condition).Where("status = ?", true)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result restaurantmodel.Restaurant

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (s *sqlStore) FindRestaurantWithoutStatus(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db

	db = db.Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result restaurantmodel.Restaurant

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
