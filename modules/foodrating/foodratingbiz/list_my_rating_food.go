package foodratingbiz

// type listUserRatingFood struct {
// 	foodRatingStore foodratingstore.FoodRatingStore
// }

// func NewlistUserRatingFood(foodRatingStore foodratingstore.FoodRatingStore) *listUserRatingFood {
// 	return &listUserRatingFood{
// 		foodRatingStore: foodRatingStore,
// 	}
// }

// func (biz *listUserRatingFood) ListRatingFoodBiz(ctx context.Context, foodId int) ([]foodratingmodel.FoodRating, error) {
// 	var paging common.Paging
// 	paging.Fulfill()
// 	result, err := biz.foodRatingStore.ListFoodRating(ctx, map[string]interface{}{"food_id": foodId}, nil, &paging, "User")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }
