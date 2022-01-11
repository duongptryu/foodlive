package restaurantstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurant/restaurantmodel"
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

	if v := filter; v != nil {
		//if v.CityId > 0 {
		//	db = db.Where("city_id = ?", v.CityId)
		//}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if paging.FakeCursor > 0 {
		db = db.Where("id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
