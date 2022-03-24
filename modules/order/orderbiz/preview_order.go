package orderbiz

import (
	"context"
	"foodlive/component/paymentprovider"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/ordermodel"
)

type previewOrderBiz struct {
	cartStore cartstore.CartStore
}

func NewPreviewOrder(cartStore cartstore.CartStore) *previewOrderBiz {
	return &previewOrderBiz{
		cartStore: cartStore,
	}
}

func (biz *previewOrderBiz) PreviewOrderBiz(ctx context.Context, userId int, rinkebyProvider paymentprovider.CryptoPaymentProvider) (*ordermodel.PreviewOrder, error) {
	cartFilter := cartmodel.Filter{}
	listCart, err := biz.cartStore.ListCartItem(ctx, map[string]interface{}{"user_id": userId}, &cartFilter, "Food")
	if err != nil {
		return nil, err
	}

	//generate order
	var totalPrice float64
	var foods []ordermodel.FoodQuantity
	for i, _ := range listCart {
		foods = append(foods, ordermodel.FoodQuantity{
			listCart[i].Food,
			listCart[i].Quantity,
		})
		totalPrice += listCart[i].Food.Price * float64(listCart[i].Quantity)
	}

	priceEth, err := rinkebyProvider.ParsePriceToEth(ctx, totalPrice)

	result := ordermodel.PreviewOrder{
		Foods:         foods,
		ShipFee:       10000,
		TotalPrice:    totalPrice,
		TotalPriceEth: priceEth,
	}

	return &result, nil
}
