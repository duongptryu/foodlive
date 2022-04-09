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

type userUpdateRatingFoodBiz struct {
	foodStore       foodstore.FoodStore
	foodRatingStore foodratingstore.FoodRatingStore
	pubSub          pubsub.PubSub
}

func NewUserUpdateRatingFoodBiz(foodStore foodstore.FoodStore,
	foodRatingStore foodratingstore.FoodRatingStore,
	pubSub pubsub.PubSub) *userUpdateRatingFoodBiz {
	return &userUpdateRatingFoodBiz{
		foodStore:       foodStore,
		foodRatingStore: foodRatingStore,
		pubSub:          pubSub,
	}
}
func (biz *userUpdateRatingFoodBiz) UpdateRestaurantRatingBiz(ctx context.Context, id int, userId int, data *foodratingmodel.FoodRatingUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	rstDb, err := biz.foodRatingStore.FindFoodRating(ctx, map[string]interface{}{"id": id, "user_id": userId})
	if err != nil {
		return err
	}
	if rstDb.Id == 0 {
		return common.ErrDataNotFound(foodratingmodel.EntityName)
	}

	if err := biz.foodRatingStore.UpdateFoodRating(ctx, id, data); err != nil {
		return err
	}

	//pubsub to calculate rating of restaurant
	err = biz.pubSub.Publish(ctx, common.TopicUserCreateRestaurantRating, pubsub.NewMessage(&foodratingmodel.FoodRatingCreate{
		FoodId: rstDb.FoodId,
	}))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
