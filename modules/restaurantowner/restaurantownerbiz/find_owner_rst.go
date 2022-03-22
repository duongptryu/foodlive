package restaurantownerbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
)

type findOwnerRst struct {
	store restaurantownerstore.RestaurantOwnerStore
}

func NewFindOwnerRst(store restaurantownerstore.RestaurantOwnerStore) *findOwnerRst {
	return &findOwnerRst{store: store}
}

func (biz *findOwnerRst) FindOwnerRst(ctx context.Context, userId int) (*restaurantownermodel.OwnerRestaurant, error) {

	result, err := biz.store.FindOwnerRestaurant(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}
	if result.Id == 0 {
		return nil, common.ErrDataNotFound(restaurantownermodel.EntityName)
	}
	return result, nil
}
