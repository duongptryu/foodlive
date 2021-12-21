package foodstore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/food/foodmodel"
)

func (s *sqlStore) ListFood(ctx context.Context,
	condition map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]foodmodel.Food, error) {
	var result []foodmodel.Food

	db := s.db

	db = db.Table(foodmodel.Food{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.CategoryId > 0 {
			db = db.Where("category_id = ?", v.CategoryId)
		}
		if v.Name != "" {
			db = db.Where("name = ?", v.Name)
		}
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
