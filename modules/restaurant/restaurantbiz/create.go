package restaurantbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	createRestaurantStore CreateRestaurantStore
}

func NewCreateRestaurantBiz(createRestaurantStore CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{
		createRestaurantStore: createRestaurantStore,
	}
}

func (biz *createRestaurantBiz) CreateRestaurantBiz(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.createRestaurantStore.CreateRestaurant(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
