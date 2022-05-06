package orderbiz

import (
	"context"
	"fmt"
	"foodlive/common"
	"foodlive/component/paymentprovider"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/useraddress/useraddressstore"
	"math"
)

type previewOrderBiz struct {
	cartStore        cartstore.CartStore
	userAddressStore useraddressstore.UserAddressStore
	restaurantStore  restaurantstore.RestaurantStore
}

func NewPreviewOrder(cartStore cartstore.CartStore, restaurantStore restaurantstore.RestaurantStore, userAddressStore useraddressstore.UserAddressStore) *previewOrderBiz {
	return &previewOrderBiz{
		cartStore:        cartStore,
		restaurantStore:  restaurantStore,
		userAddressStore: userAddressStore,
	}
}

func (biz *previewOrderBiz) PreviewOrderBiz(ctx context.Context, userId int, data *ordermodel.OrderPreviewReq, rinkebyProvider paymentprovider.CryptoPaymentProvider) (*ordermodel.PreviewOrder, error) {
	cartFilter := cartmodel.Filter{}

	userAddress, err := biz.userAddressStore.FindUserAddressById(ctx, map[string]interface{}{"user_id": userId, "id": data.AddressId})
	if err != nil {
		return nil, err
	}

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
	distance := common.Distance(userAddress.Lat, userAddress.Lng, rst.Lat, rst.Lng, "K")

	//generate order
	var totalPrice float64
	var foods []ordermodel.FoodQuantity
	for i := range listCart {
		foods = append(foods, ordermodel.FoodQuantity{
			listCart[i].Food,
			listCart[i].Quantity,
		})
		totalPrice += listCart[i].Food.Price * float64(listCart[i].Quantity)
	}

	shipFee := int(rst.ShippingFeePerKm * math.Round(distance))

	totalPrice += float64(shipFee)

	priceEth, err := rinkebyProvider.ParsePriceToEth(ctx, totalPrice/23000)
	if err != nil {
		return nil, err
	}
	newPriceEth := fmt.Sprintf("%.18f", priceEth)

	result := ordermodel.PreviewOrder{
		Foods:         foods,
		ShipFee:       shipFee,
		Distance:      distance,
		TotalPrice:    totalPrice,
		TotalPriceEth: newPriceEth,
	}

	return &result, nil
}
