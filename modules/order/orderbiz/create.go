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
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/useraddress/useraddressstore"
	log "github.com/sirupsen/logrus"
	"time"
)

type createOrderBiz struct {
	orderStore       orderstore.OrderStore
	cartStore        cartstore.CartStore
	restaurantStore  restaurantstore.RestaurantStore
	orderDetailStore orderdetailstore.OrderDetailStore
	orderTracking    ordertrackingstore.OrderTrackingStore
	userAddressStore useraddressstore.UserAddressStore
	paymentProvider  paymentprovider.PaymentProvider
}

func NewCreateOrderBiz(orderStore orderstore.OrderStore, orderDetailStore orderdetailstore.OrderDetailStore,
	orderTracking ordertrackingstore.OrderTrackingStore, userAddressStore useraddressstore.UserAddressStore, cartStore cartstore.CartStore, paymentProvider paymentprovider.PaymentProvider, restaurantStore restaurantstore.RestaurantStore) *createOrderBiz {
	return &createOrderBiz{
		orderStore:       orderStore,
		orderTracking:    orderTracking,
		orderDetailStore: orderDetailStore,
		cartStore:        cartStore,
		userAddressStore: userAddressStore,
		paymentProvider:  paymentProvider,
		restaurantStore:  restaurantStore,
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

	if len(listCart) == 0 {
		return nil, ordermodel.ErrCartEmpty
	}

	//get lat lng of rst
	rst, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": listCart[0].RestaurantId}, nil)
	if err != nil {
		return nil, err
	}

	//calculate distance
	distance := common.Distance(addressDb.Lat, addressDb.Lng, rst.Lat, rst.Lng, "K")

	//generate order
	var totalPrice float64
	for i, _ := range listCart {
		totalPrice += listCart[i].Food.Price * float64(listCart[i].Quantity)
	}

	shipFee := int(rst.ShippingFeePerKm * distance)

	totalPrice += float64(shipFee)

	if totalPrice > float64(50000) {
		return nil, ordermodel.ErrMoneyTooBig
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

	orderCreateRevert, err := biz.orderStore.CreateOrder(ctx, &order)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	checkoutResp, err := biz.paymentProvider.SendRequestPayment(ctx, &order, "Payment for food delivery service and ship to "+addressDb.Addr)
	if err != nil {
		orderCreateRevert.Rollback()
		return nil, err
	}
	if checkoutResp.ResultCode != 0 {
		//fail
		orderCreateRevert.Rollback()
		return nil, ordermodel.ErrPaymentFailed
	}

	//success
	//create order detail
	var orderDetails = make([]orderdetailmodel.OrderDetailCreate, len(listCart))
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

	createOrderDetailRevert, err := biz.orderDetailStore.CreateBulkOrderDetail(ctx, orderDetails)
	if err != nil {
		orderCreateRevert.Rollback()
		return nil, err
	}

	//create order tracking
	orderTracking := ordertrackingmodel.OrderTrackingCreate{
		OrderId: order.Id,
		State:   ordertrackingmodel.StateWaitingPayment,
	}

	orderTrackingRevert, err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking)
	if err != nil {
		createOrderDetailRevert.Rollback()
		orderCreateRevert.Rollback()
		return nil, err
	}

	if err = biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId}); err != nil {
		orderCreateRevert.Rollback()
		createOrderDetailRevert.Rollback()
		orderTrackingRevert.Rollback()
		return nil, err
	}

	orderCreateRevert.Commit()
	orderTrackingRevert.Commit()
	createOrderDetailRevert.Commit()

	go biz.cancelOrder(ctx, order.Id, 10)

	return checkoutResp, nil
}

func (biz *createOrderBiz) CreateOrderCryptoBiz(ctx context.Context, userId int, data *ordermodel.Checkout) (*ordermodel.PaymentCoinResp, error) {
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

	//get lat lng of rst
	rst, err := biz.restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": listCart[0].RestaurantId}, nil)
	if err != nil {
		return nil, err
	}

	//calculate distance
	distance := common.Distance(addressDb.Lat, addressDb.Lng, rst.Lat, rst.Lng, "K")

	//generate order
	var totalPrice float64
	for i, _ := range listCart {
		totalPrice += listCart[i].Food.Price * float64(listCart[i].Quantity)
	}

	shipFee := int(rst.ShippingFeePerKm * distance)

	totalPrice += float64(shipFee)

	var order = ordermodel.OrderCreate{
		UserId:         userId,
		RestaurantId:   listCart[0].RestaurantId,
		TotalPrice:     totalPrice,
		ShipperId:      1,
		Status:         true,
		UserAddressOri: addressDb.Addr,
		TypePayment:    ordermodel.TypeCrypto,
	}

	orderCreateRevert, err := biz.orderStore.CreateOrder(ctx, &order)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	//create order detail
	var orderDetails = make([]orderdetailmodel.OrderDetailCreate, len(listCart))
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

	orderDetailRevert, err := biz.orderDetailStore.CreateBulkOrderDetail(ctx, orderDetails)
	if err != nil {
		orderCreateRevert.Rollback()
		return nil, err
	}

	//create order tracking
	orderTracking := ordertrackingmodel.OrderTrackingCreate{
		OrderId: order.Id,
		State:   ordertrackingmodel.StateWaitingPayment,
	}

	orderTrackingRevert, err := biz.orderTracking.CreateOrderTracking(ctx, &orderTracking)
	if err != nil {
		orderCreateRevert.Rollback()
		orderDetailRevert.Rollback()
		return nil, err
	}
	if err = biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId}); err != nil {
		orderCreateRevert.Rollback()
		orderDetailRevert.Rollback()
		orderTrackingRevert.Rollback()
		log.Error(err)
	}

	orderCreateRevert.Commit()
	orderDetailRevert.Commit()
	orderTrackingRevert.Commit()

	result := ordermodel.PaymentCoinResp{
		OrderId: order.Id,
		Web:     fmt.Sprintf("https://foodlive.tech/order/%v", order.Id),
		App:     fmt.Sprintf("https://metamask.app.link/dapp/foodlive.tech/order/%v", order.Id),
	}

	go biz.cancelOrder(ctx, order.Id, 5)

	return &result, nil
}

func (biz *createOrderBiz) cancelOrder(ctx context.Context, orderId int, minute uint) {
	time.AfterFunc(time.Duration(minute)*time.Minute, func() {
		order, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId})
		if err != nil {
			log.Error(err)
			return
		}
		if order.Id == 0 {
			log.Error(common.ErrDataNotFound(ordermodel.EntityName))
			return
		}

		if order.Status == false {
			return
		}

		orderTracking, err := biz.orderTracking.FindOrderTracking(ctx, map[string]interface{}{"order_id": orderId})
		if err != nil {
			log.Error(err)
			return
		}

		if orderTracking.Id == 0 {
			log.Error(common.ErrDataNotFound(ordertrackingmodel.EntityName))
			return
		}

		if orderTracking.State != ordertrackingmodel.StateWaitingPayment {
			return
		}

		status := false

		orderUpdate := ordermodel.OrderUpdate{
			Status: &status,
		}

		if err := biz.orderStore.UpdateOrder(ctx, orderId, &orderUpdate); err != nil {
			log.Error(err)
			return
		}

		orderTrackingUpdate := ordertrackingmodel.OrderTrackingUpdate{
			State: ordertrackingmodel.StateCancel,
		}
		if err := biz.orderTracking.UpdateOrderTracking(ctx, orderId, &orderTrackingUpdate); err != nil {
			log.Error(err)
			return
		}
	})
}
