package cartbiz

import (
	"context"
	"foodlive/modules/cart/cartmodel"
)

type listFoodInCartBiz struct {
	cartStore CartStore
}

func NewListFoodInCartBiz(cartStore CartStore) *listFoodInCartBiz {
	return &listFoodInCartBiz{
		cartStore: cartStore,
	}
}

func (biz *listFoodInCartBiz) ListFoodInCartBiz(ctx context.Context, userId int, filter *cartmodel.Filter) ([]cartmodel.CartItem, error) {
	result, err := biz.cartStore.ListCartItem(ctx, map[string]interface{}{"user_id": userId}, filter, "Food")
	if err != nil {
		return nil, err
	}
	return result, nil
}
