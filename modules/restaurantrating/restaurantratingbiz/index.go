package restaurantratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

type RestaurantRatingStore interface {
	CreateRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error
	DeleteRestaurantRating(ctx context.Context, id int) error
	ListRestaurantRating(ctx context.Context, condition map[string]interface{}, filter *restaurantratingmodel.Filter, paging *common.Paging, moreKey ...string) ([]restaurantratingmodel.RestaurantRating, error)
	CalculateAVGPoint(ctx context.Context,
		condition map[string]interface{},
	) (float64, error)
	UpdateRestaurantRating(ctx context.Context, id int, data *restaurantratingmodel.RestaurantRatingUpdate) error
	FindRestaurantRating(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantratingmodel.RestaurantRating, error)
}

type RestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, filter *restaurantmodel.Filter, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
