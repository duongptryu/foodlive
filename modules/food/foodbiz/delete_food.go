package foodbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/food/foodmodel"
	"fooddelivery/modules/restaurant/restaurantmodel"
)

type deleteFoodBiz struct {
	foodStore       FoodStore
	restaurantStore RestaurantStore
}

func NewDeleteFoodBiz(foodStore FoodStore, restaurantStore RestaurantStore) *deleteFoodBiz {
	return &deleteFoodBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
	}
}

func (biz *deleteFoodBiz) DeleteFoodBiz(ctx context.Context, id int, userId int) error {
	//check if restaurant exist
	foodDb, err := biz.foodStore.FindFood(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if foodDb.Id == 0 {
		return foodmodel.ErrFoodNotFound
	}

	//check if restaurant exist
	rDb, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": foodDb.RestaurantId})
	if err != nil {
		return err
	}

	if rDb.Id == 0 {
		return restaurantmodel.ErrRestaurantNotFound
	}

	if rDb.OwnerId != userId {
		return common.ErrPermissionDenied
	}

	if err := biz.foodStore.DeleteFood(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	return nil
}
