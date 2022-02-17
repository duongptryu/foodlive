package categorystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
