package orderbiz

import (
	"context"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"log"
	"strconv"
)

type webHookPaymentBiz struct {
	orderStore         orderstore.OrderStore
	orderTrackingStore ordertrackingstore.OrderTrackingStore
}

func NewWebHookPaymentBiz(orderStore orderstore.OrderStore, orderTrackingStore ordertrackingstore.OrderTrackingStore) *webHookPaymentBiz {
	return &webHookPaymentBiz{
		orderStore:         orderStore,
		orderTrackingStore: orderTrackingStore,
	}
}

func (biz *webHookPaymentBiz) WebHookPaymentBiz(ctx context.Context, data *ordermodel.WebHookPayment) error {
	if err := data.Validate(); err != nil {
		return err
	}

	orderId, _ := strconv.Atoi(data.OrderID)
	var state string

	if data.ErrorCode != "0" {
		state = ordertrackingmodel.StatePaymentFail
	} else {
		state = ordertrackingmodel.StatePreparing
	}

	orderTracking := ordertrackingmodel.OrderTrackingUpdate{
		State:  state,
		Status: true,
	}
	if err := biz.orderTrackingStore.UpdateOrder(ctx, orderId, &orderTracking); err != nil {
		log.Println(err)
		return err
	}

	//payment success, create order tracking and order detail and notify user
	return nil
}
