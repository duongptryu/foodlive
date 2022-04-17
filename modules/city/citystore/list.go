package citystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/city/citymodel"
)

func (s *sqlStore) ListCity(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) ([]citymodel.City, error) {
	var result []citymodel.City

	db := s.db

	db = db.Table(citymodel.City{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
