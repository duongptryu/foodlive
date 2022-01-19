package cartbiz

import (
	"context"
	"foodlive/modules/cart/cartmodel"
)

type deleteFoodInCartBiz struct {
	cartStore CartStore
}

func NewDeleteFoodInCartBiz(cartStore CartStore) *deleteFoodInCartBiz {
	return &deleteFoodInCartBiz{
		cartStore: cartStore,
	}
}

func (biz *deleteFoodInCartBiz) DeleteAFoodInCartBiz(ctx context.Context, userId, foodId int) error {

	itemDb, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": userId, "food_id": foodId})
	if err != nil {
		return err
	}
	if itemDb.UserId == 0 {
		return cartmodel.ErrItemDoesNotExist
	}

	//cart of user is empty, add item to cart
	if err := biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId, "food_id": foodId}); err != nil {
		return err
	}
	return nil
}

func (biz *deleteFoodInCartBiz) DeleteAllFoodInCartBiz(ctx context.Context, userId int) error {
	itemDb, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}
	if itemDb.UserId == 0 {
		return cartmodel.ErrItemDoesNotExist
	}

	//cart of user is empty, add item to cart
	if err := biz.cartStore.DeleteCartItem(ctx, map[string]interface{}{"user_id": userId}); err != nil {
		return err
	}
	return nil
}
