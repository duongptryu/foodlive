package rstcategorystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
)

func (s *sqlStore) ListRestaurantByCategory(ctx context.Context,
	condition map[string]interface{},
	filter *rstcategorymodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]rstcategorymodel.RstCategory, error) {
	var result []rstcategorymodel.RstCategory

	db := s.db

	db = db.Table(rstcategorymodel.RstCategory{}.TableName()).Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := filter; v != nil {
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
