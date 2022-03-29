package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListRestaurant(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
		if v.OwnerId != 0 {
			db = db.Where("owner_id = ?", v.OwnerId)
		}
		if v.Lng != 0 && v.Lat != 0 {
			db = db.Select("*, (6371 * acos(cos(radians(?)) * cos(radians(lat)) * cos(radians(lng) - radians(?)) + sin(radians(?)) * sin(radians(lat)))) AS distance", v.Lat, v.Lng, v.Lat).Having("distance < 2").Order("distance")
		} else {
			if paging.FakeCursor > 0 {
				db = db.Where("id < ?", paging.FakeCursor).Order("id desc")
			} else {
				db = db.Offset((paging.Page - 1) * paging.Limit).Order("id desc")
			}
		}
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
