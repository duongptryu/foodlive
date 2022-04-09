package foodlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
	"foodlive/modules/foodlike/foodlikestore"
	"foodlive/pubsub"
)

type userUnlikeFoodBiz struct {
	store  foodlikestore.FoodLikeStore
	pubSub pubsub.PubSub
}

func NewUserUnlikeFoodBiz(store foodlikestore.FoodLikeStore, pubSub pubsub.PubSub) *userUnlikeFoodBiz {
	return &userUnlikeFoodBiz{
		store:  store,
		pubSub: pubSub,
	}
}

func (biz *userUnlikeFoodBiz) UserUnlikeFood(ctx context.Context, userId, foodId int) error {
	exist, err := biz.store.FindUserLikeFood(ctx, map[string]interface{}{"user_id": userId, "food_id": foodId})
	if err != nil {
		return err
	}
	if exist.UserId == 0 {
		return foodlikemodel.ErrUserNotLikeFoodYet
	}

	err = biz.store.DeleteFoodLike(ctx, map[string]interface{}{"user_id": userId, "food_id": foodId})

	if err != nil {
		return err
	}

	biz.pubSub.Publish(ctx, common.TopicUserUnlikeFood, pubsub.NewMessage(&foodlikemodel.FoodLike{
		FoodId: foodId,
		UserId:       userId,
	}))

	return nil
}
