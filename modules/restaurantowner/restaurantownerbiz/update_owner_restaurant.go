package restaurantownerbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
)

type updateOwnerRstBiz struct {
	store restaurantownerstore.RestaurantOwnerStore
}

func NewUpdateOwnerRstBiz(store restaurantownerstore.RestaurantOwnerStore) *updateOwnerRstBiz {
	return &updateOwnerRstBiz{
		store: store,
	}
}

func (biz *updateOwnerRstBiz) UpdateOwnerRstBiz(ctx context.Context, id int, data *restaurantownermodel.OwnerRestaurantUpdate) error {
	if err := biz.store.UpdateOwnerRestaurant(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantownermodel.EntityName, err)
	}
	return nil
}
