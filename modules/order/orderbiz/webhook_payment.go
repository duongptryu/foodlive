package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"log"
	"strconv"
)

type webHookPaymentBiz struct {
	orderStore         orderstore.OrderStore
	orderTrackingStore ordertrackingstore.OrderStore
}

func NewWebHookPaymentBiz(orderStore orderstore.OrderStore, orderTrackingStore ordertrackingstore.OrderStore) *webHookPaymentBiz {
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
		state = ordermodel.PaymentFailStatus
	} else {
		state = ordermodel.PrepareStatus
	}

	orderTracking := ordertrackingmodel.OrderTrackingUpdate{
		common.SQLModelUpdate{},
		state,
		true,
	}
	if err := biz.orderTrackingStore.UpdateOrder(ctx, orderId, &orderTracking); err != nil {
		log.Println(err)
		return err
	}

	//payment success, create order tracking and order detail and notify user
	return nil
}
