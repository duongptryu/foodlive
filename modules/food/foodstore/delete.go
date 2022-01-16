package foodstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodmodel"
)

func (s *sqlStore) DeleteFood(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
