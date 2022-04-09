package rstcategorystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
)

func (s *sqlStore) CreateRestaurantCategory(ctx context.Context, data []rstcategorymodel.RstCategoryCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
