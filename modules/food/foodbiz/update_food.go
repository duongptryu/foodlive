package foodbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/food/foodmodel"
	"fooddelivery/modules/restaurant/restaurantmodel"
)

type updateFoodBiz struct {
	foodStore       FoodStore
	restaurantStore RestaurantStore
}

func NewUpdateFoodBiz(foodStore FoodStore, restaurantStore RestaurantStore) *updateFoodBiz {
	return &updateFoodBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
	}
}

func (biz *updateFoodBiz) UpdateFoodBiz(ctx context.Context, id int, userId int, data *foodmodel.FoodUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

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

	if foodDb.CategoryId != data.CategoryId {
		//check if category is exist
	}

	if err := biz.foodStore.UpdateFood(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(foodmodel.EntityName, err)
	}

	return nil
}
