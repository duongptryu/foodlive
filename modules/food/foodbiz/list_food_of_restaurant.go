package foodbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/restaurant/restaurantmodel"
)

type listFoodOfRestaurantBiz struct {
	foodStore       FoodStore
	restaurantStore RestaurantStore
}

func NewListFoodOfRestaurantBiz(foodStore FoodStore, restaurantStore RestaurantStore) *listFoodOfRestaurantBiz {
	return &listFoodOfRestaurantBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
	}
}

func (biz *listFoodOfRestaurantBiz) ListFoodOfRestaurantBiz(ctx context.Context, restaurantId int, paging *common.Paging, filter *foodmodel.Filter) ([]foodmodel.Food, error) {
	//check if restaurant exist
	rst, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": restaurantId})
	if err != nil {
		return nil, err
	}

	if rst.Id == 0 {
		return nil, restaurantmodel.ErrRestaurantNotFound
	}

	result, err := biz.foodStore.ListFood(ctx, map[string]interface{}{"restaurant_id": restaurantId}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listFoodOfRestaurantBiz) UserListFoodOfRestaurantBiz(ctx context.Context, restaurantId int, paging *common.Paging, filter *foodmodel.Filter) ([]foodmodel.Food, error) {
	//check if restaurant exist
	rst, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": restaurantId})
	if err != nil {
		return nil, err
	}

	if rst.Id == 0 {
		return nil, restaurantmodel.ErrRestaurantNotFound
	}

	result, err := biz.foodStore.ListFood(ctx, map[string]interface{}{"restaurant_id": restaurantId, "status": true}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
