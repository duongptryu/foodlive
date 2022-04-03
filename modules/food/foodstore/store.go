package foodstore

import (
	"context"
	"foodlive/modules/food/foodmodel"
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

type FoodStore interface {
	ListFoodWithoutPaging(ctx context.Context,
		condition map[string]interface{},
		filter *foodmodel.Filter,
		moreKey ...string,
	) ([]foodmodel.Food, error)
}
