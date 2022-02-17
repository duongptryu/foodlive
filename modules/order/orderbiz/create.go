package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type createOrderBiz struct {
	orderStore orderstore.OrderStore
}

func NewCreateOrderBiz(orderStore orderstore.OrderStore) *createOrderBiz {
	return &createOrderBiz{
		orderStore: orderStore,
	}
}

func (biz *createOrderBiz) CreateOrderBiz(ctx context.Context, data *ordermodel.OrderCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.orderStore.CreateOrder(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	return nil
}
