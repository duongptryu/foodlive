package cartbiz

import (
	"context"
	"foodlive/modules/cart/cartmodel"
)

type addFoodToCartBiz struct {
	cartStore CartStore
}

func NewAddFoodToCartBiz(cartStore CartStore) *addFoodToCartBiz {
	return &addFoodToCartBiz{
		cartStore: cartStore,
	}
}

func (biz *addFoodToCartBiz) AddFoodToCartBiz(ctx context.Context, data *cartmodel.CartItemCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	itemDb, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": data.UserId}, "Food")
	if err != nil {
		return err
	}
	if itemDb.UserId != 0 {
		//cart of user already exist 1 or more item
		if itemDb.Food.Id != data.FoodId {
			return cartmodel.ErrFoodInAnotherRestaurant
		}

		exist, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": data.UserId, "food_id": data.FoodId})
		if err != nil {
			return err
		}
		if exist.UserId != 0 {
			return cartmodel.ErrItemAlreadyExist
		}
	}

	//cart of user is empty, add item to cart
	if err := biz.cartStore.CreateCartItem(ctx, data); err != nil {
		return err
	}
	return nil
}
