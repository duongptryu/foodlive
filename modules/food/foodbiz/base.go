package foodbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/restaurant/restaurantmodel"
)

type FoodStore interface {
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
	DeleteFood(ctx context.Context, id int) error
	FindFood(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error)
	ListFood(ctx context.Context, condition map[string]interface{}, filter *foodmodel.Filter, paging *common.Paging, moreKey ...string) ([]foodmodel.Food, error)
	UpdateFood(ctx context.Context, id int, data *foodmodel.FoodUpdate) error
}

type RestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
