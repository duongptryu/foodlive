package orderbiz

import (
	"context"
	"fmt"
	"foodlive/common"
	"foodlive/component/paymentprovider"
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

func (biz *findOrderBiz) FindOrderBiz(ctx context.Context, orderId int, userId int, rinkebyProvider paymentprovider.CryptoPaymentProvider) (*ordermodel.OrderResponse, error) {
	order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId, "user_id": userId}, "Restaurant")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	if order.Id == 0 {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	if order.UserId != userId {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	orderDetail, err := biz.orderDetail.FindOrderDetail(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return nil, err
	}

	priceEth, err := rinkebyProvider.ParsePriceToEth(ctx, order.TotalPrice/23000)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	newPriceEth := fmt.Sprintf("%.18f", priceEth)

	order.TotalPriceEth = newPriceEth

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

func (biz *findOrderBiz) FindOrderOfRestaurantBiz(ctx context.Context, orderId int, userId int, rinkebyProvider paymentprovider.CryptoPaymentProvider) (*ordermodel.OrderResponse, error) {
	order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId}, "Restaurant")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	if order.Id == 0 {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	if order.Restaurant.OwnerId != userId {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	if order.UserId != userId {
		return nil, common.ErrDataNotFound(ordermodel.EntityName)
	}

	orderDetail, err := biz.orderDetail.FindOrderDetail(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return nil, err
	}

	priceEth, err := rinkebyProvider.ParsePriceToEth(ctx, order.TotalPrice/23000)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	newPriceEth := fmt.Sprintf("%.18f", priceEth)

	order.TotalPriceEth = newPriceEth

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

func (biz *findOrderBiz) FindOrderCryptoBiz(ctx context.Context, orderId int, rinkebyProvider paymentprovider.CryptoPaymentProvider) (*ordermodel.OrderResponse, error) {
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

	priceEth, err := rinkebyProvider.ParsePriceToEth(ctx, order.TotalPrice/23000)

	newPriceEth := fmt.Sprintf("%.18f", priceEth)

	order.TotalPriceEth = newPriceEth

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
