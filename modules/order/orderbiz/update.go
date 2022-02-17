package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type updateOrderBiz struct {
	orderStore orderstore.OrderStore
}

func NewUpdateOrderBiz(orderStore orderstore.OrderStore) *updateOrderBiz {
	return &updateOrderBiz{
		orderStore: orderStore,
	}
}

func (biz *updateOrderBiz) UpdateOrderBiz(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	//check if restaurant exist
	orderDb, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if orderDb.Id == 0 {
		return common.ErrDataNotFound(ordermodel.EntityName)
	}

	if err := biz.orderStore.UpdateOrder(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(ordermodel.EntityName, err)
	}

	return nil
}
