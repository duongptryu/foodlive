package restaurantownerstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
)

func (s *sqlStore) FindOwnerRestaurant(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*restaurantownermodel.OwnerRestaurant, error) {
	db := s.db
	var result restaurantownermodel.OwnerRestaurant

	db = db.Where(conditions)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
