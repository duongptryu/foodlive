package categorystore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/category/categorymodel"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
