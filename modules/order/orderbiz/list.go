package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type listOrderBiz struct {
	orderStore orderstore.OrderStore
}

func NewListOrderBiz(orderStore orderstore.OrderStore) *listOrderBiz {
	return &listOrderBiz{
		orderStore: orderStore,
	}
}

func (biz *listOrderBiz) ListOrderBiz(ctx context.Context, userId int, paging *common.Paging, filter *ordermodel.Filter) ([]ordermodel.Order, error) {
	result, err := biz.orderStore.ListOrder(ctx, map[string]interface{}{"user_id": userId}, filter, paging, "Restaurant", "OrderTracking")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}

func (biz *listOrderBiz) ListMyCurrentOrderBiz(ctx context.Context, userId int, paging *common.Paging, filter *ordermodel.Filter) ([]ordermodel.Order, error) {
	result, err := biz.orderStore.ListOrder(ctx, map[string]interface{}{"user_id": userId, "status": true}, filter, paging, "Restaurant", "OrderTracking")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}
