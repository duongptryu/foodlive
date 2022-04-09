package foodratingbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodrating/foodratingmodel"
	"foodlive/modules/foodrating/foodratingstore"
	"foodlive/pubsub"
	"log"
)

type userRatingFoodBiz struct {
	foodStore       foodstore.FoodStore
	foodRatingStore foodratingstore.FoodRatingStore
	pubSub          pubsub.PubSub
}

func NewUserRatingFoodBiz(foodStore foodstore.FoodStore,
	foodRatingStore foodratingstore.FoodRatingStore,
	pubSub pubsub.PubSub) *userRatingFoodBiz {
	return &userRatingFoodBiz{
		foodStore: foodStore,
		foodRatingStore: foodRatingStore,
		pubSub:                pubSub,
	}
}

func (biz *userRatingFoodBiz) UserRatingFoodBiz(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	exist, err := biz.foodRatingStore.FindFoodRating(ctx, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})
	if err != nil {
		return err
	}
	if exist.Id != 0 {
		return foodratingmodel.ErrAlreadyRating
	}

	rstDb, err := biz.foodStore.FindFood(ctx, map[string]interface{}{"id": data.FoodId})
	if err != nil {
		return err
	}
	if rstDb.Id == 0 {
		return common.ErrDataNotFound(foodratingmodel.EntityName)
	}

	if err := biz.foodRatingStore.CreateFoodRating(ctx, data); err != nil {
		return err
	}

	//pubsub to calculate rating of restaurant
	err = biz.pubSub.Publish(ctx, common.TopicUserRatingFood, pubsub.NewMessage(data))
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
