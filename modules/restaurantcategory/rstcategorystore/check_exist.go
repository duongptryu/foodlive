package rstcategorystore

import (
	"context"
	"foodlive/common"
)

func (s *sqlStore) CountRstCategory(ctx context.Context, condition map[string]interface{}) (int64, error) {
	db := s.db

	var result int64

	db = db.Where(condition)

	if err := db.Count(&result).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	return result, nil
}
