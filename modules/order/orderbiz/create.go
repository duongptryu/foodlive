package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/paymentprovider"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type createOrderBiz struct {
	orderStore      orderstore.OrderStore
	cartStore       cartstore.CartStore
	paymentProvider paymentprovider.PaymentProvider
}

func NewCreateOrderBiz(orderStore orderstore.OrderStore, cartStore cartstore.CartStore, paymentProvider paymentprovider.PaymentProvider) *createOrderBiz {
	return &createOrderBiz{
		orderStore:      orderStore,
		cartStore:       cartStore,
		paymentProvider: paymentProvider,
	}
}

func (biz *createOrderBiz) CreateOrderBiz(ctx context.Context, userId int, data *ordermodel.OrderCreate) (*paymentprovider.TransactionResp, error) {
	cartFilter := cartmodel.Filter{}
	listCart, err := biz.cartStore.ListCartItem(ctx, map[string]interface{}{"user_id": userId}, &cartFilter, "Food")
	if err != nil {
		return nil, err
	}

	//generate order
	var totalPrice float64
	for i, _ := range listCart {
		totalPrice += listCart[i].Food.Price * float64(listCart[i].Quantity)
	}

	var order = ordermodel.OrderCreate{
		common.SQLModelCreate{},
		userId,
		totalPrice,
		1,
		true,
	}

	if err := biz.orderStore.CreateOrder(ctx, &order); err != nil {
		return nil, common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	checkoutResp, err := biz.paymentProvider.SendRequestPayment(ctx, data, "Test")
	if err != nil {
		return nil, err
	}

	return checkoutResp, nil
}
