package cartbiz

import (
	"context"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/restaurant/restaurantmodel"
)

type CartStore interface {
	CreateCartItem(ctx context.Context, data *cartmodel.CartItemCreate) error
	DeleteCartItem(ctx context.Context, conditions map[string]interface{}) error
	FindCartItem(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*cartmodel.CartItem, error)
	UpdateCartItem(ctx context.Context, data *cartmodel.CartItemUpdate) error
	ListCartItem(ctx context.Context, condition map[string]interface{}, filter *cartmodel.Filter, moreKey ...string) ([]cartmodel.CartItem, error)
}

type FoodStore interface {
	FindFood(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error)
}

type RestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
