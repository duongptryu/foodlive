package restaurantbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	createRestaurantStore CreateRestaurantStore
	cityStore             CityStore
}

func NewCreateRestaurantBiz(createRestaurantStore CreateRestaurantStore, cityStore CityStore) *createRestaurantBiz {
	return &createRestaurantBiz{
		createRestaurantStore: createRestaurantStore,
		cityStore:             cityStore,
	}
}

func (biz *createRestaurantBiz) CreateRestaurantBiz(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	cate, err := biz.cityStore.FindCity(ctx, map[string]interface{}{"id": data.CityId, "status": true})
	if err != nil {
		return err
	}

	if cate.Id == 0 {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, common.ErrDataNotFound(categorymodel.EntityName))
	}

	if err := biz.createRestaurantStore.CreateRestaurant(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
