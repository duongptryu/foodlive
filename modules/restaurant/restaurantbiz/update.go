package restaurantbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
	UpdateRestaurantStatus(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdateStatus) error
}

type updateRestaurantBiz struct {
	Store     UpdateRestaurantStore
	cityStore CityStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore, cityStore CityStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store, cityStore}
}

func (biz *updateRestaurantBiz) UpdateRestaurantBiz(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.Store.FindRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Id == 0 || oldData.Status == false {
		return restaurantmodel.ErrRestaurantNotFound
	}

	if data.CityId != oldData.CityId {
		cate, err := biz.cityStore.FindCity(ctx, map[string]interface{}{"id": data.CityId, "status": true})
		if err != nil {
			return err
		}

		if cate.Id == 0 {
			return common.ErrCannotCreateEntity(restaurantmodel.EntityName, common.ErrDataNotFound(categorymodel.EntityName))
		}
	}

	if err := biz.Store.UpdateRestaurant(ctx, id, data); err != nil {
		return err
	}
	return nil
}

func (biz *updateRestaurantBiz) UpdateRestaurantStatusBiz(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdateStatus) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.Store.FindRestaurant(ctx, map[string]interface{}{"id": id, "status": false})
	if err != nil {
		return err
	}
	if oldData.Id == 0 {
		return restaurantmodel.ErrRestaurantNotFound
	}

	if oldData.Status == *data.Status {
		return restaurantmodel.ErrStatusAlreadySet
	}

	if err := biz.Store.UpdateRestaurantStatus(ctx, id, data); err != nil {
		return err
	}
	return nil
}
