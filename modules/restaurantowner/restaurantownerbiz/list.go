package restaurantownerbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
)

type listOwnerRestaurant struct {
	store restaurantownerstore.RestaurantOwnerStore
}

func NewListOwnerRestaurant(store restaurantownerstore.RestaurantOwnerStore) *listOwnerRestaurant {
	return &listOwnerRestaurant{store: store}
}

func (biz *listOwnerRestaurant) ListOwnerRestaurant(ctx context.Context, filter *restaurantownermodel.Filter,
	paging *common.Paging) ([]restaurantownermodel.OwnerRestaurant, error) {

	result, err := biz.store.ListOwnerRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
