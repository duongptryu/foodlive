package restaurantratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) ListRestaurantRating(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantratingmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantratingmodel.RestaurantRating, error) {
	var result []restaurantratingmodel.RestaurantRating

	db := s.db

	db = db.Table(restaurantratingmodel.RestaurantRating{}.TableName()).Where(condition)

	//if v := filter; v != nil {
	//}

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

func (s *sqlStore) CountRestaurantRating(ctx context.Context,
	condition map[string]interface{},
) (int, error) {
	db := s.db

	db = db.Table(restaurantratingmodel.RestaurantRating{}.TableName()).Where(condition)

	var result int64
	if err := db.Count(&result).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	return int(result), nil
}
