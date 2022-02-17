package foodbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/restaurant/restaurantmodel"
)

type createFoodBiz struct {
	foodStore       FoodStore
	restaurantStore RestaurantStore
	categoryStore   CategoryStore
}

func NewCreateFoodBiz(foodStore FoodStore, restaurantStore RestaurantStore, categoryStore CategoryStore) *createFoodBiz {
	return &createFoodBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
		categoryStore:   categoryStore,
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
	cateDb, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"id": data.CategoryId, "status": true})
	if err != nil {
		return err
	}

	if cateDb.Id == 0 {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, common.ErrDataNotFound(categorymodel.EntityName))
	}

	if err := biz.foodStore.CreateFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	return nil
}
