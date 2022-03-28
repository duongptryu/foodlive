package orderbiz

import (
	"context"
	"foodlive/modules/historypayment/historypaymentstore"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
	log "github.com/sirupsen/logrus"
	"time"
)

type confirmPaymentCryptoBiz struct {
	orderStore          orderstore.OrderStore
	historyPaymentStore historypaymentstore.HistoryPaymentStore
	orderTrackingStore  ordertrackingstore.OrderTrackingStore
}

func NewConfirmPaymentCryptoBiz(orderStore orderstore.OrderStore, orderTrackingStore ordertrackingstore.OrderTrackingStore, historyPaymentStore historypaymentstore.HistoryPaymentStore) *confirmPaymentCryptoBiz {
	return &confirmPaymentCryptoBiz{
		orderStore:          orderStore,
		historyPaymentStore: historyPaymentStore,
		orderTrackingStore:  orderTrackingStore,
	}
}

func (biz *confirmPaymentCryptoBiz) ConfirmCryptoPayment(ctx context.Context, data *ordermodel.PaymentOrderEvent) error {
	//create order history
	//paymentHistory := historypaymentmodel.HistoryPaymentCreate{
	//	Amount:       data.Value,
	//	OrderID:      data.OrderID,
	//	TransID:      data.TransID,
	//	Message:      data.Message,
	//	RequestID:    data.RequestID,
	//	ResponseTime: data.ResponseTime,
	//	ExtraData:    data.ExtraData,
	//	ResultCode:   data.ResultCode,
	//	LocalMessage: data.LocalMessage,
	//	PayType:      data.PayType,
	//}

	//if err := biz.historyPaymentStore.CreateHistoryPayment(ctx, &paymentHistory); err != nil {
	//	log.Error(err)
	//}

	orderUpdate := ordermodel.OrderUpdate{
		TotalPriceEth: data.Value,
		TxnHash:       data.Hash,
	}

	if err := biz.orderStore.UpdateOrder(ctx, data.OrderId, &orderUpdate); err != nil {
		log.Error(err)
		return err
	}

	orderTracking := ordertrackingmodel.OrderTrackingUpdate{
		State: ordertrackingmodel.StatePreparing,
	}
	if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, data.OrderId, &orderTracking); err != nil {
		log.Error(err)
		return err
	}

	//payment success, create order tracking and order detail and notify user
	go func(biz *confirmPaymentCryptoBiz) {
		time.AfterFunc(60*time.Second, func() {
			//update status order tracking
			orderTrackingOnTheWay := ordertrackingmodel.OrderTrackingUpdate{
				State: ordertrackingmodel.StateOnTheWay,
			}
			if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, data.OrderId, &orderTrackingOnTheWay); err != nil {
				log.Println(err)
				return
			}

			time.AfterFunc(60*time.Second, func() {
				orderTrackingDelivered := ordertrackingmodel.OrderTrackingUpdate{
					State: ordertrackingmodel.StateDelivered,
				}
				if err := biz.orderTrackingStore.UpdateOrderTracking(ctx, data.OrderId, &orderTrackingDelivered); err != nil {
					log.Println(err)
					return
				}
				status := false

				orderUpdate := ordermodel.OrderUpdate{Status: &status}
				if err := biz.orderStore.UpdateOrder(ctx, data.OrderId, &orderUpdate); err != nil {
					log.Println(err)
					return
				}
			})
		})
	}(biz)

	return nil
}
