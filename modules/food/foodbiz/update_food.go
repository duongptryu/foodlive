package foodbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/restaurant/restaurantmodel"
)

type updateFoodBiz struct {
	foodStore       foodstore.FoodStore
	restaurantStore RestaurantStore
	categoryStore   CategoryStore
}

func NewUpdateFoodBiz(foodStore foodstore.FoodStore, restaurantStore RestaurantStore, categoryStore CategoryStore) *updateFoodBiz {
	return &updateFoodBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
		categoryStore:   categoryStore,
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
		cateDb, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"id": data.CategoryId, "status": true})
		if err != nil {
			return err
		}

		if cateDb.Id == 0 {
			return common.ErrCannotCreateEntity(foodmodel.EntityName, common.ErrDataNotFound(categorymodel.EntityName))
		}
	}

	if err := biz.foodStore.UpdateFood(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(foodmodel.EntityName, err)
	}

	return nil
}
