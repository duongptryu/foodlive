package restaurantlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/pubsub"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type userUnLikeRestaurantBiz struct {
	store  UserUnLikeRestaurantStore
	pubSub pubsub.PubSub
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, pubSub pubsub.PubSub) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store:  store,
		pubSub: pubSub,
	}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotUnLikeRestaurant(err)
	}

	biz.pubSub.Publish(ctx, common.TopicUserDisLikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		restaurantId,
		userId,
		nil,
		nil,
	}))

	return nil
}
