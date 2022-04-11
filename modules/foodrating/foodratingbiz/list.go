package foodratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodrating/foodratingmodel"
	"foodlive/modules/foodrating/foodratingstore"
)

type listUserRatingFood struct {
	foodRatingStore foodratingstore.FoodRatingStore
}

func NewlistUserRatingFood(foodRatingStore foodratingstore.FoodRatingStore) *listUserRatingFood {
	return &listUserRatingFood{
		foodRatingStore: foodRatingStore,
	}
}

func (biz *listUserRatingFood) ListRatingFoodBiz(ctx context.Context, foodId int, filter *foodratingmodel.Filter, paging *common.Paging) ([]foodratingmodel.FoodRating, error) {
	result, err := biz.foodRatingStore.ListFoodRating(ctx, map[string]interface{}{"food_id": foodId}, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (biz *listUserRatingFood) LisMytRatingFoodBiz(ctx context.Context, userId int, filter *foodratingmodel.Filter, paging *common.Paging) ([]foodratingmodel.FoodRating, error) {
	result, err := biz.foodRatingStore.ListFoodRating(ctx, map[string]interface{}{"user_id": userId}, filter, paging, "User", "Food")
	if err != nil {
		return nil, err
	}

	return result, nil
}
