package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
)

type updateOrderBiz struct {
	orderStore         orderstore.OrderStore
	orderTrackingStore ordertrackingstore.OrderTrackingStore
}

func NewUpdateOrderBiz(orderStore orderstore.OrderStore, orderTrackingStore ordertrackingstore.OrderTrackingStore) *updateOrderBiz {
	return &updateOrderBiz{
		orderStore:         orderStore,
		orderTrackingStore: orderTrackingStore,
	}
}

func (biz *updateOrderBiz) RstConfirmPrepareDone(ctx context.Context, orderId, ownerId int) error {
	order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId}, "Restaurant")
	if err != nil {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if order.Id == 0 {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if order.Status == false {
		return ordermodel.ErrOrderExpire
	}
	if order.Restaurant.OwnerId != ownerId {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	orderTracking, err := biz.orderTrackingStore.FindOrderTracking(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if orderTracking.Id == 0 {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if orderTracking.State != ordertrackingmodel.StatePreparing {
		return ordermodel.ErrOrderExpire
	}

	orderTrackingUpdate := ordertrackingmodel.OrderTrackingUpdate{
		State: ordertrackingmodel.StateOnTheWay,
	}

	if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, orderId, &orderTrackingUpdate); err != nil {
		return err
	}

	return nil
}

func (biz *updateOrderBiz) UserConfirmReceived(ctx context.Context, orderId, userId int) error {
	order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId, "user_id": userId})
	if err != nil {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if order.Id == 0 {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if order.Status == false {
		return ordermodel.ErrOrderExpire
	}

	orderTracking, err := biz.orderTrackingStore.FindOrderTracking(ctx, map[string]interface{}{"order_id": orderId})
	if err != nil {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if orderTracking.Id == 0 {
		return common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	if orderTracking.State != ordertrackingmodel.StateOnTheWay {
		return ordermodel.ErrOrderExpire
	}

	orderTrackingUpdate := ordertrackingmodel.OrderTrackingUpdate{
		State: ordertrackingmodel.StateDelivered,
	}
	if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, orderId, &orderTrackingUpdate); err != nil {
		return err
	}

	status := false
	orderUpdate := ordermodel.OrderUpdate{
		Status: &status,
	}
	if err := biz.orderStore.UpdateOrder(ctx, orderId, &orderUpdate); err != nil {
		return err
	}

	return nil
}
