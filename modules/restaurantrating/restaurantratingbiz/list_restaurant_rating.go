package restaurantratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
)

type listRestaurantRatingBiz struct {
	restaurantStore       RestaurantStore
	restaurantRatingStore RestaurantRatingStore
}

func NewListRestaurantRatingBiz(restaurantStore RestaurantStore, restaurantRatingStore RestaurantRatingStore) *listRestaurantRatingBiz {
	return &listRestaurantRatingBiz{
		restaurantStore:       restaurantStore,
		restaurantRatingStore: restaurantRatingStore,
	}
}

func (biz *listRestaurantRatingBiz) ListRestaurantRatingBiz(ctx context.Context, rstId int, paging *common.Paging, filter *restaurantratingmodel.Filter) ([]restaurantratingmodel.RestaurantRating, error) {
	rstDb, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": rstId})
	if err != nil {
		return nil, err
	}
	if rstDb.Id == 0 {
		return nil, common.ErrDataNotFound(restaurantlikemodel.EntityName)
	}

	result, err := biz.restaurantRatingStore.ListRestaurantRating(ctx, map[string]interface{}{"restaurant_id": rstId}, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	return result, nil
}
