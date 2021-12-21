package foodbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/food/foodmodel"
	"fooddelivery/modules/restaurant/restaurantmodel"
)

type createFoodBiz struct {
	foodStore       FoodStore
	restaurantStore RestaurantStore
}

func NewCreateFoodBiz(foodStore FoodStore, restaurantStore RestaurantStore) *createFoodBiz {
	return &createFoodBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
	}
}

func (biz *createFoodBiz) CreateFoodBiz(ctx context.Context, data *foodmodel.FoodCreate, userId int) error {
	if err := data.Validate(); err != nil {
		return err
	}

	//check if restaurant exist
	result, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": data.RestaurantId})
	if err != nil {
		return err
	}

	if result.Id == 0 {
		return restaurantmodel.ErrRestaurantNotFound
	}

	if result.OwnerId != userId {
		return restaurantmodel.ErrRestaurantNotFound
	}

	//check if category is exist

	if err := biz.foodStore.CreateFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	return nil
}
