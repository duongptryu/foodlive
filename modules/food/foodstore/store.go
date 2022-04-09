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
	IncreaseLikeCount(ctx context.Context, id int) error
	DecreaseLikeCount(ctx context.Context, id int) error
	UpdateRating(ctx context.Context, foodId int, rating float64) error
	FindFood(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error) 
}
