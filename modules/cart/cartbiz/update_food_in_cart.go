package cartbiz

import (
	"context"
	"foodlive/modules/cart/cartmodel"
)

type updateFoodInCartBiz struct {
	cartStore CartStore
}

func NewUpdateFoodInCartBiz(cartStore CartStore) *updateFoodInCartBiz {
	return &updateFoodInCartBiz{
		cartStore: cartStore,
	}
}

func (biz *updateFoodInCartBiz) UpdateFoodInCartBiz(ctx context.Context, data *cartmodel.CartItemUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	itemDb, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": data.UserId, "food_id": data.FoodId})
	if err != nil {
		return err
	}
	if itemDb.UserId == 0 {
		return cartmodel.ErrItemDoesNotExist
	}

	if err := biz.cartStore.UpdateCartItem(ctx, data); err != nil {
		return err
	}
	return nil
}
