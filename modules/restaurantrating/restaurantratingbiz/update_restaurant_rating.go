package restaurantratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
	"foodlive/pubsub"
)

type updateRestaurantRatingBiz struct {
	restaurantStore       RestaurantStore
	restaurantRatingStore RestaurantRatingStore
	pubSub                pubsub.PubSub
}

func UpdateRestaurantRatingBiz(restaurantStore RestaurantStore, restaurantRatingStore RestaurantRatingStore, pubSub pubsub.PubSub) *updateRestaurantRatingBiz {
	return &updateRestaurantRatingBiz{
		restaurantStore:       restaurantStore,
		restaurantRatingStore: restaurantRatingStore,
		pubSub:                pubSub,
	}
}

func (biz *updateRestaurantRatingBiz) UpdateRestaurantRatingBiz(ctx context.Context, id int, userId int, data *restaurantratingmodel.RestaurantRatingUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	rstDb, err := biz.restaurantRatingStore.FindRestaurantRating(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if rstDb.Id == 0 {
		return common.ErrDataNotFound(restaurantratingmodel.EntityName)
	}

	if rstDb.UserId != userId {
		return common.ErrPermissionDenied
	}

	if err := biz.restaurantRatingStore.UpdateRestaurantRating(ctx, id, data); err != nil {
		return err
	}

	//pubsub to calculate rating of restaurant
	return nil
}
