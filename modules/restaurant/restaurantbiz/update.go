package restaurantbiz

import (
	"context"
	"fooddelivery/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
	UpdateRestaurantStatus(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdateStatus) error
}

type updateRestaurantBiz struct {
	Store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store}
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
