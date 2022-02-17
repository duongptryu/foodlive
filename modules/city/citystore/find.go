package citystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/city/citymodel"
)

func (s *sqlStore) FindCity(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*citymodel.City, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result citymodel.City

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
