package restaurantratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

type listMyRatingBiz struct {
	restaurantStore       RestaurantStore
	restaurantRatingStore RestaurantRatingStore
}

func NewListMyRatingBiz(restaurantStore RestaurantStore, restaurantRatingStore RestaurantRatingStore) *listRestaurantRatingBiz {
	return &listRestaurantRatingBiz{
		restaurantStore:       restaurantStore,
		restaurantRatingStore: restaurantRatingStore,
	}
}

func (biz *listRestaurantRatingBiz) ListMyRatingBiz(ctx context.Context, paging *common.Paging, filter *restaurantratingmodel.Filter) ([]restaurantratingmodel.RestaurantRating, error) {
	var result []restaurantratingmodel.RestaurantRating
	var err error
	if filter.RstId != 0 {
		result, err = biz.restaurantRatingStore.ListRestaurantRating(ctx, map[string]interface{}{"restaurant_id": filter.RstId, "user_id": filter.UserId}, filter, paging, "User")
		if err != nil {
			return nil, err
		}
	} else {
		result, err = biz.restaurantRatingStore.ListRestaurantRating(ctx, map[string]interface{}{"user_id": filter.UserId}, filter, paging, "User")
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
