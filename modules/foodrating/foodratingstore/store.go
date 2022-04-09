package foodratingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodrating/foodratingmodel"

	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type FoodRatingStore interface {
	CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error
	DeleteFoodRating(ctx context.Context, condition map[string]interface{}) error
	FindFoodRating(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*foodratingmodel.FoodRating, error) 
	ListFoodRating(ctx context.Context,
		condition map[string]interface{},
		filter *foodratingmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]foodratingmodel.FoodRating, error)
	CalculateAVGPoint(ctx context.Context,
		condition map[string]interface{},
	) (float64, error)
	UpdateFoodRating(ctx context.Context, id int, data *foodratingmodel.FoodRatingUpdate) error 
}
