package foodbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodlike/foodlikestore"
	"foodlive/modules/restaurant/restaurantmodel"
	log "github.com/sirupsen/logrus"
)

type listFoodOfRestaurantBiz struct {
	foodStore       foodstore.FoodStore
	restaurantStore RestaurantStore
	foodLikeStore   foodlikestore.FoodLikeStore
}

func NewListFoodOfRestaurantBiz(foodStore foodstore.FoodStore, restaurantStore RestaurantStore, foodLikeStore foodlikestore.FoodLikeStore) *listFoodOfRestaurantBiz {
	return &listFoodOfRestaurantBiz{
		foodStore:       foodStore,
		restaurantStore: restaurantStore,
		foodLikeStore:   foodLikeStore,
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

func (biz *listFoodOfRestaurantBiz) UserListFoodOfRestaurantBiz(ctx context.Context, restaurantId int, userId int, paging *common.Paging, filter *foodmodel.Filter) ([]foodmodel.Food, error) {
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

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.foodLikeStore.GetFoodLiked(ctx, ids, userId)
	if err != nil {
		log.Println("Cannot get food liked: ", err)
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			if v, exist := mapResLike[item.Id]; exist {
				result[i].IsLike = v
			}
		}
	}

	return result, nil
}

func (biz *listFoodOfRestaurantBiz) ListAllFood(ctx context.Context, userId int, paging *common.Paging, filter *foodmodel.Filter) ([]foodmodel.Food, error) {
	result, err := biz.foodStore.ListFood(ctx, map[string]interface{}{"status": true}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.foodLikeStore.GetFoodLiked(ctx, ids, userId)
	if err != nil {
		log.Println("Cannot get food liked: ", err)
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			if v, exist := mapResLike[item.Id]; exist {
				result[i].IsLike = v
			}
		}
	}

	return result, nil
}
