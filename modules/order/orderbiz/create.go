package orderbiz

import (
	"context"
	"fmt"
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
	log "github.com/sirupsen/logrus"
)

type createOrderBiz struct {
	orderStore       orderstore.OrderStore
	cartStore        cartstore.CartStore
	orderDetailStore orderdetailstore.OrderDetailStore
	orderTracking    ordertrackingstore.OrderTrackingStore
	userAddressStore useraddressstore.UserAddressStore
	paymentProvider  paymentprovider.PaymentProvider
}

func NewCreateOrderBiz(orderStore orderstore.OrderStore, orderDetailStore orderdetailstore.OrderDetailStore,
	orderTracking ordertrackingstore.OrderTrackingStore, userAddressStore useraddressstore.UserAddressStore, cartStore cartstore.CartStore, paymentProvider paymentprovider.PaymentProvider) *createOrderBiz {
	return &createOrderBiz{
		orderStore:       orderStore,
		orderTracking:    orderTracking,
		orderDetailStore: orderDetailStore,
		cartStore:        cartStore,
		userAddressStore: userAddressStore,
		paymentProvider:  paymentProvider,
	}
}

func (biz *createOrderBiz) CreateOrderMomoBiz(ctx context.Context, userId int, data *ordermodel.Checkout) (*paymentprovider.TransactionResp, error) {
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
		TypePayment:    ordermodel.TypeMomo,
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
			State:   ordertrackingmodel.StatePaymentFail,
		}
		if err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking); err != nil {
			log.Println(err)
			return nil, err
		}

		return nil, ordermodel.ErrPaymentFailed
	}

	go func() {
		//create order detail
		var orderDetails []orderdetailmodel.OrderDetailCreate
		for i, v := range listCart {
			orderDetails[i] = orderdetailmodel.OrderDetailCreate{
				OrderId: order.Id,
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
			State:   ordertrackingmodel.StateWaitingPayment,
		}
		if err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking); err != nil {
			log.Error(err)
			return
		}
		if err = biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId}); err != nil {
			log.Error(err)
		}
	}()

	return checkoutResp, nil
}

func (biz *createOrderBiz) CreateOrderCryptoBiz(ctx context.Context, userId int, data *ordermodel.Checkout, rinkebyProvider paymentprovider.CryptoPaymentProvider) (*ordermodel.PaymentCoinResp, error) {
	addressDb, err := biz.userAddressStore.FindUserAddressById(ctx, map[string]interface{}{"user_id": userId, "id": data.UserAddrId})
	if err != nil {
		return nil, err
	}

	cartFilter := cartmodel.Filter{}
	listCart, err := biz.cartStore.ListCartItem(ctx, map[string]interface{}{"user_id": userId}, &cartFilter, "Food")
	if err != nil {
		return nil, err
	}

	if len(listCart) == 0 {
		return nil, ordermodel.ErrCartEmpty
	}

	//generate order
	var totalPrice float64
	for i, _ := range listCart {
		totalPrice += listCart[i].Food.Price * float64(listCart[i].Quantity)
	}

	priceEth, err := rinkebyProvider.ParsePriceToEth(ctx, totalPrice)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	var order = ordermodel.OrderCreate{
		UserId:         userId,
		RestaurantId:   listCart[0].RestaurantId,
		TotalPrice:     totalPrice,
		ShipperId:      1,
		Status:         true,
		UserAddressOri: addressDb.Addr,
		TypePayment:    ordermodel.TypeCrypto,
		TotalPriceEth:  priceEth,
	}

	if err := biz.orderStore.CreateOrder(ctx, &order); err != nil {
		return nil, common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	go func(cartItems []cartmodel.CartItem) {
		//create order detail
		var orderDetails = make([]orderdetailmodel.OrderDetailCreate, len(cartItems))
		for i, v := range cartItems {
			orderDetails[i] = orderdetailmodel.OrderDetailCreate{
				OrderId: order.Id,
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
			State:   ordertrackingmodel.StateWaitingPayment,
		}
		if err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking); err != nil {
			log.Println(err)
			return
		}
		if err = biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId}); err != nil {
			log.Error(err)
		}
	}(listCart)

	result := ordermodel.PaymentCoinResp{
		OrderId: order.Id,
		Web:     fmt.Sprintf("https://foodlive.tech/order/%v", order.Id),
		App:     fmt.Sprintf("https://metamask.app.link/dapp/foodlive.tech/order/%v", order.Id),
	}

	return &result, nil
}
