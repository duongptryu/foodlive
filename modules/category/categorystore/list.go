package categorystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
)

func (s *sqlStore) ListCategory(ctx context.Context,
	condition map[string]interface{},
	filter *categorymodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]categorymodel.Category, error) {
	var result []categorymodel.Category

	db := s.db

	db = db.Table(categorymodel.Category{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("name LIKE ?", "%"+v.Name+"%")
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
