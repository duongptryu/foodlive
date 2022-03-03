package orderbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/paymentprovider"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/orderdetail/orderdetailmodel"
	"foodlive/modules/orderdetail/orderdetailstore"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"foodlive/modules/useraddress/useraddressstore"
	"log"
)

type createOrderBiz struct {
	orderStore       orderstore.OrderStore
	cartStore        cartstore.CartStore
	orderDetailStore orderdetailstore.OrderStore
	orderTracking    ordertrackingstore.OrderStore
	userAddressStore useraddressstore.UserAddressStore
	paymentProvider  paymentprovider.PaymentProvider
}

func NewCreateOrderBiz(orderStore orderstore.OrderStore, orderDetailStore orderdetailstore.OrderStore,
	orderTracking ordertrackingstore.OrderStore, userAddressStore useraddressstore.UserAddressStore, cartStore cartstore.CartStore, paymentProvider paymentprovider.PaymentProvider) *createOrderBiz {
	return &createOrderBiz{
		orderStore:       orderStore,
		orderTracking:    orderTracking,
		orderDetailStore: orderDetailStore,
		cartStore:        cartStore,
		userAddressStore: userAddressStore,
		paymentProvider:  paymentProvider,
	}
}

func (biz *createOrderBiz) CreateOrderBiz(ctx context.Context, userId int, data *ordermodel.Checkout) (*paymentprovider.TransactionResp, error) {
	addressDb, err := biz.userAddressStore.FindUserAddressById(ctx, map[string]interface{}{"user_id": userId, "id": data.UserAddrId})
	if err != nil {
		return nil, err
	}

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
		UserId:         userId,
		RestaurantId:   listCart[0].RestaurantId,
		TotalPrice:     totalPrice,
		ShipperId:      1,
		Status:         true,
		UserAddressOri: addressDb.Addr,
	}

	if err := biz.orderStore.CreateOrder(ctx, &order); err != nil {
		return nil, common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	checkoutResp, err := biz.paymentProvider.SendRequestPayment(ctx, &order, "Payment for food delivery service and ship to "+addressDb.Addr)
	if err != nil {
		return nil, err
	}
	if checkoutResp.ErrorCode != 0 {
		//Update order, order detail, order tracking
		//create order tracking
		orderTracking := ordertrackingmodel.OrderTrackingCreate{
			OrderId: order.Id,
			State:   ordermodel.PaymentFailStatus,
			Status:  true,
		}
		if err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking); err != nil {
			log.Println(err)
			return nil, err
		}

		return nil, ordermodel.ErrPaymentFailed
	}

	go func() {
		biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId})

		//create order detail
		var orderDetails []orderdetailmodel.OrderDetailCreate
		for i, v := range listCart {
			orderDetails[i] = orderdetailmodel.OrderDetailCreate{
				UserId:       userId,
				OrderId:      order.Id,
				RestaurantId: v.RestaurantId,
				FoodOrigin: &common.FoodOrigin{
					Id:     v.Food.Id,
					Name:   v.Food.Name,
					Price:  v.Food.Price,
					Images: v.Food.Images,
				},
				Price:    v.Food.Price,
				Quantity: v.Quantity,
			}
		}

		err = biz.orderDetailStore.CreateBulkOrderDetail(ctx, orderDetails)
		if err != nil {
			log.Println(err)
			return
		}

		//create order tracking
		orderTracking := ordertrackingmodel.OrderTrackingCreate{
			OrderId: order.Id,
			State:   ordermodel.PaymentStatus,
			Status:  true,
		}
		if err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking); err != nil {
			log.Println(err)
			return
		}
	}()

	return checkoutResp, nil
}
