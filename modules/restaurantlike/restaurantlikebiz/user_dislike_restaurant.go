package restaurantlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/pubsub"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
	FindUserLikeRst(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantlikemodel.Like, error)
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
	exist, err := biz.store.FindUserLikeRst(ctx, map[string]interface{}{"user_id": userId, "restaurant_id": restaurantId})
	if err != nil {
		return err
	}
	if exist.UserId == 0 {
		return restaurantlikemodel.ErrUserCannotUnLikeRestaurant(nil)
	}

	err = biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotUnLikeRestaurant(err)
	}

	biz.pubSub.Publish(ctx, common.TopicUserDisLikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		RestaurantId: restaurantId,
		UserId:       userId,
	}))

	return nil
}
