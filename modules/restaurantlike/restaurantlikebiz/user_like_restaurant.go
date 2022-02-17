package restaurantlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store  UserLikeRestaurantStore
	pubSub pubsub.PubSub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	pubSub pubsub.PubSub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:  store,
		pubSub: pubSub,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotLikeRestaurant(err)
	}

	biz.pubSub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	return nil
}
