package restaurantlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	FindUserLikeRst(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantlikemodel.Like, error)
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
	exist, err := biz.store.FindUserLikeRst(ctx, map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId})
	if err != nil {
		return err
	}
	if exist.UserId != 0 {
		return restaurantlikemodel.ErrUserAlreadyLikeRestaurant
	}

	err = biz.store.Create(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrUserCannotLikeRestaurant(err)
	}

	err = biz.pubSub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
