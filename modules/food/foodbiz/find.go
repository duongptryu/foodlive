package foodbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodlike/foodlikestore"
)

type findFoodBiz struct {
	foodStore       foodstore.FoodStore
	restaurantStore RestaurantStore
	foodLikeStore   foodlikestore.FoodLikeStore
}

func NewFindFoodBiz(foodStore foodstore.FoodStore, restaurantStore RestaurantStore, foodLikeStore foodlikestore.FoodLikeStore) *findFoodBiz {
	return &findFoodBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
		foodLikeStore:   foodLikeStore,
	}
}

func (biz *findFoodBiz) FindFoodBiz(ctx context.Context, foodId int, userId int) (*foodmodel.Food, error) {
	//check if restaurant exist

	result, err := biz.foodStore.FindFood(ctx, map[string]interface{}{"id": foodId}, "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	isLike, err := biz.foodLikeStore.GetFoodLiked(ctx, []int{result.Id}, userId)
	if err != nil {
		return nil, err
	}

	result.IsLike = isLike[result.Id]

	return result, nil
}
