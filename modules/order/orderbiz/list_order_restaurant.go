package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/restaurant/restaurantmodel"
	"foodlive/modules/restaurant/restaurantstore"
)

type listOrderRestaurantBiz struct {
	restaurantStore restaurantstore.RestaurantStore
	orderStore      orderstore.OrderStore
}

func NewListOrderRestaurantBiz(orderStore orderstore.OrderStore, restaurantStore restaurantstore.RestaurantStore) *listOrderRestaurantBiz {
	return &listOrderRestaurantBiz{
		orderStore:      orderStore,
		restaurantStore: restaurantStore,
	}
}

func (biz *listOrderRestaurantBiz) ListOrderRestaurantBiz(ctx context.Context, ownerId int, rstId int, paging *common.Paging, filter *ordermodel.Filter) ([]ordermodel.Order, error) {
	rst, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"owner_id": ownerId, "id": rstId})
	if err != nil {
		return nil, err
	}
	if rst.Id == 0 {
		return nil, common.ErrDataNotFound(restaurantmodel.EntityName)
	}

	result, err := biz.orderStore.ListOrder(ctx, map[string]interface{}{"restaurant_id": rstId}, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}

func (biz *listOrderRestaurantBiz) ListCurrentOrderRestaurantBiz(ctx context.Context, ownerId int, rstId int, paging *common.Paging, filter *ordermodel.Filter) ([]ordermodel.Order, error) {
	rst, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"owner_id": ownerId, "id": rstId})
	if err != nil {
		return nil, err
	}
	if rst.Id == 0 {
		return nil, common.ErrDataNotFound(restaurantmodel.EntityName)
	}

	result, err := biz.orderStore.ListOrder(ctx, map[string]interface{}{"restaurant_id": rstId, "status": true}, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}
