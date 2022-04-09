package foodlikestore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

type FoodLikeStore interface {
	Create(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error
	DeleteFoodLike(ctx context.Context, condition map[string]interface{}) error
	FindUserLikeFood(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*foodlikemodel.FoodLike, error)
	ListUsersLikFood(ctx context.Context, conditions map[string]interface{}, filter *foodlikemodel.Filter, paging *common.Paging, moreKeys ...string) ([]foodlikemodel.FoodLike, error)
}
