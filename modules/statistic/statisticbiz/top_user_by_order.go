package statisticbiz

import (
	"context"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type statsUserByOrderBiz struct {
	orderStore orderstore.OrderStore
}

func NewStatsUserByOrderBiz(orderStore orderstore.OrderStore) *statsUserByOrderBiz {
	return &statsUserByOrderBiz{
		orderStore: orderStore,
	}
}

func (biz *statsUserByOrderBiz) StatsUserByOrderBiz(ctx context.Context) ([]ordermodel.OrderGroupByUser, error) {
	result, err := biz.orderStore.ListOrderGroupByUser(ctx, nil, "User")
	if err != nil {
		return nil, err
	}
	return result, nil
}
