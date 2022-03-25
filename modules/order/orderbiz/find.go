package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/orderdetail/orderdetailstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
)

type findOrderBiz struct {
	orderStore    orderstore.OrderStore
	orderDetail   orderdetailstore.OrderDetailStore
	orderTracking ordertrackingstore.OrderTrackingStore
}

func NewFindOrderBiz(orderStore orderstore.OrderStore, orderDetail orderdetailstore.OrderDetailStore,
	orderTracking ordertrackingstore.OrderTrackingStore) *findOrderBiz {
	return &findOrderBiz{
		orderStore:    orderStore,
		orderTracking: orderTracking,
		orderDetail:   orderDetail,
	}
}

func (biz *findOrderBiz) FindOrderBiz(ctx context.Context, orderId int, userId int) (*ordermodel.OrderResponse, error) {
	order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId, "user_id": userId}, "Restaurant")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	if order.Id == 0 {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	orderDetail, err := biz.orderDetail.FindOrderDetail(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return nil, err
	}

	orderTracking, err := biz.orderTracking.FindOrderTracking(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return nil, err
	}

	resp := ordermodel.OrderResponse{
		Order:         order,
		OrderDetail:   orderDetail,
		OrderTracking: orderTracking,
	}

	return &resp, nil
}

func (biz *findOrderBiz) FindOrderCryptoBiz(ctx context.Context, orderId int) (*ordermodel.OrderResponse, error) {
	order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId}, "Restaurant")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	if order.Id == 0 {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	if order.Status == false {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}
	if order.TypePayment != ordermodel.TypeCrypto {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	orderDetail, err := biz.orderDetail.FindOrderDetail(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return nil, err
	}

	orderTracking, err := biz.orderTracking.FindOrderTracking(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return nil, err
	}

	if orderTracking.State != ordertrackingmodel.StateWaitingPayment {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	resp := ordermodel.OrderResponse{
		Order:         order,
		OrderDetail:   orderDetail,
		OrderTracking: orderTracking,
	}

	return &resp, nil
}

