package restaurantownerstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
)

func (s *sqlStore) ListOwnerRestaurant(ctx context.Context, condition map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]restaurantownermodel.OwnerRestaurant, error) {
	db := s.db
	var result []restaurantownermodel.OwnerRestaurant

	db = db.Table(restaurantownermodel.OwnerRestaurant{}.TableName()).Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
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
