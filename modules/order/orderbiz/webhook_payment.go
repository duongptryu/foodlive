package orderbiz

import (
	"context"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type webHookPaymentBiz struct {
	orderStore orderstore.OrderStore
}

func NewWebHookPaymentBiz(orderStore orderstore.OrderStore) *webHookPaymentBiz {
	return &webHookPaymentBiz{
		orderStore: orderStore,
	}
}

func (biz *webHookPaymentBiz) WebHookPaymentBiz (ctx context.Context, data *ordermodel.WebHookPayment) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if data.ErrorCode != "0" {
		// payment fail, update order and notify user
		return nil
	}

	//payment success, create order tracking and order detail and notify user
	return nil
}