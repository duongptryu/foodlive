package restaurantbiz

import (
	"context"
	"foodlive/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id int) error
}

type deleteRestaurant struct {
	Store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurant {
	return &deleteRestaurant{store}
}

func (biz *deleteRestaurant) DeleteRestaurantBiz(ctx context.Context, id int) error {
	oldData, err := biz.Store.FindRestaurant(ctx, map[string]interface{}{"id": id, "status": true})
	if err != nil {
		return err
	}
	if oldData.Id == 0 {
		return restaurantmodel.ErrRestaurantNotFound
	}

	if err := biz.Store.DeleteRestaurant(ctx, id); err != nil {
		return err
	}
	return nil
}
