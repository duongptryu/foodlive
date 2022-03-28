package orderbiz

import (
	"context"
	"foodlive/modules/historypayment/historypaymentmodel"
	"foodlive/modules/historypayment/historypaymentstore"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type webHookPaymentBiz struct {
	orderStore          orderstore.OrderStore
	historyPaymentStore historypaymentstore.HistoryPaymentStore
	orderTrackingStore  ordertrackingstore.OrderTrackingStore
}

func NewWebHookPaymentBiz(orderStore orderstore.OrderStore, orderTrackingStore ordertrackingstore.OrderTrackingStore, historyPaymentStore historypaymentstore.HistoryPaymentStore) *webHookPaymentBiz {
	return &webHookPaymentBiz{
		orderStore:          orderStore,
		historyPaymentStore: historyPaymentStore,
		orderTrackingStore:  orderTrackingStore,
	}
}

func (biz *webHookPaymentBiz) WebHookPaymentBiz(ctx context.Context, data *ordermodel.WebHookPayment) error {
	//create order history
	paymentHistory := historypaymentmodel.HistoryPaymentCreate{
		Amount:       data.Amount,
		OrderID:      data.OrderID,
		OrderInfo:    data.OrderInfo,
		OrderType:    data.OrderType,
		TransID:      data.TransID,
		Message:      data.Message,
		RequestID:    data.RequestID,
		ResponseTime: data.ResponseTime,
		ExtraData:    data.ExtraData,
		ResultCode:   data.ResultCode,
		LocalMessage: data.LocalMessage,
		PayType:      data.PayType,
	}

	if err := biz.historyPaymentStore.CreateHistoryPayment(ctx, &paymentHistory); err != nil {
		log.Error(err)
	}

	orderId, _ := strconv.Atoi(data.OrderID)
	var state string

	if data.ResultCode != 0 {
		state = ordertrackingmodel.StatePaymentFail
	} else {
		state = ordertrackingmodel.StatePreparing
	}

	orderTracking := ordertrackingmodel.OrderTrackingUpdate{
		State: state,
	}
	if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, orderId, &orderTracking); err != nil {
		log.Println(err)
		return err
	}

	//payment success, create order tracking and order detail and notify user
	go func(biz *webHookPaymentBiz) {
		time.AfterFunc(60*time.Second, func() {
			//update status order tracking
			orderTrackingOnTheWay := ordertrackingmodel.OrderTrackingUpdate{
				State: ordertrackingmodel.StateOnTheWay,
			}
			if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, orderId, &orderTrackingOnTheWay); err != nil {
				log.Println(err)
				return
			}

			time.AfterFunc(60*time.Second, func() {
				orderTrackingDelivered := ordertrackingmodel.OrderTrackingUpdate{
					State: ordertrackingmodel.StateDelivered,
				}
				if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, orderId, &orderTrackingDelivered); err != nil {
					log.Println(err)
					return
				}
				status := false

				orderUpdate := ordermodel.OrderUpdate{Status: &status}
				if err := biz.orderStore.UpdateOrder(ctx, orderId, &orderUpdate); err != nil {
					log.Println(err)
					return
				}
			})
		})
	}(biz)

	return nil
}
