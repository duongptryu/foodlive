package rstcategorystore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type RstCategoryStore interface {
	DeleteRestaurantCategory(ctx context.Context, condition map[string]interface{}) error
	CreateRestaurantCategory(ctx context.Context, data []rstcategorymodel.RstCategoryCreate) error
	CountRstCategory(ctx context.Context, condition map[string]interface{}) (int64, error)
	ListRestaurantByCategory(ctx context.Context,
		condition map[string]interface{},
		filter *rstcategorymodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]rstcategorymodel.RstCategory, error)
}
