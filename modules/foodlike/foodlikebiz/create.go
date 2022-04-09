package foodlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
	"foodlive/modules/foodlike/foodlikestore"
	"foodlive/pubsub"
	"log"
)

type userLikeFoodBiz struct {
	store  foodlikestore.FoodLikeStore
	pubSub pubsub.PubSub
}

func NewUserLikeFoodBiz(
	store foodlikestore.FoodLikeStore,
	pubSub pubsub.PubSub) *userLikeFoodBiz {
	return &userLikeFoodBiz{
		store:  store,
		pubSub: pubSub,
	}
}

func (biz *userLikeFoodBiz) UserLikeFoodBiz(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	exist, err := biz.store.FindUserLikeFood(ctx, map[string]interface{}{"user_id": data.UserId, "food_id": data.FoodId})
	if err != nil {
		return err
	}
	if exist.UserId != 0 {
		return foodlikemodel.ErrUserAlreadyLikeFood
	}

	err = biz.store.Create(ctx, data)
	if err != nil {
		return err
	}

	err = biz.pubSub.Publish(ctx, common.TopicUserLikeFood, pubsub.NewMessage(data))
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
