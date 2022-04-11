package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindRestaurant(ctx context.Context, condition map[string]interface{}, filter *restaurantmodel.Filter, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := filter; v != nil {
		if v.Lat != 0 && v.Lng != 0 {
			db = db.Select("*, (6371 * acos(cos(radians(?)) * cos(radians(lat)) * cos(radians(lng) - radians(?)) + sin(radians(?)) * sin(radians(lat)))) AS distance", v.Lat, v.Lng, v.Lat)
		}
	}

	var result restaurantmodel.Restaurant

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
