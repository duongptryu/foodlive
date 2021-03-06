package restaurantratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
	"foodlive/pubsub"
	log "github.com/sirupsen/logrus"
)

type createRestaurantRatingBiz struct {
	restaurantStore       RestaurantStore
	restaurantRatingStore RestaurantRatingStore
	pubSub                pubsub.PubSub
}

func NewCreateRestaurantRatingBiz(restaurantStore RestaurantStore, restaurantRatingStore RestaurantRatingStore, pubSub pubsub.PubSub) *createRestaurantRatingBiz {
	return &createRestaurantRatingBiz{
		restaurantStore:       restaurantStore,
		restaurantRatingStore: restaurantRatingStore,
		pubSub:                pubSub,
	}
}

func (biz *createRestaurantRatingBiz) CreateRestaurantRatingBiz(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	exist, err := biz.restaurantRatingStore.FindRestaurantRating(ctx, map[string]interface{}{"restaurant_id": data.RestaurantId, "user_id": data.UserId})
	if err != nil {
		return err
	}
	if exist.Id != 0 {
		return restaurantratingmodel.ErrAlreadyRating
	}

	rstDb, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": data.RestaurantId}, nil)
	if err != nil {
		return err
	}
	if rstDb.Id == 0 {
		return common.ErrDataNotFound(restaurantlikemodel.EntityName)
	}

	if err := biz.restaurantRatingStore.CreateRestaurantRating(ctx, data); err != nil {
		return err
	}

	//pubsub to calculate rating of restaurant
	err = biz.pubSub.Publish(ctx, common.TopicUserCreateRestaurantRating, pubsub.NewMessage(data))
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
