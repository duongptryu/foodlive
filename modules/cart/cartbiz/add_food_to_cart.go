package cartbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/food/foodmodel"
)

type addFoodToCartBiz struct {
	cartStore CartStore
	foodStore FoodStore
}

func NewAddFoodToCartBiz(cartStore CartStore, foodStore FoodStore) *addFoodToCartBiz {
	return &addFoodToCartBiz{
		cartStore: cartStore,
		foodStore: foodStore,
	}
}

func (biz *addFoodToCartBiz) AddFoodToCartBiz(ctx context.Context, data *cartmodel.CartItemCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	food, err := biz.foodStore.FindFood(ctx, map[string]interface{}{"id": data.FoodId})
	if err != nil {
		return err
	}
	if food.Id == 0 {
		return common.ErrDataNotFound(foodmodel.EntityName)
	}

	itemDb, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": data.UserId})
	if err != nil {
		return err
	}
	if itemDb.UserId != 0 {
		//cart of user already exist 1 or more item
		if itemDb.RestaurantId != food.RestaurantId {
			return cartmodel.ErrFoodInAnotherRestaurant
		}

		exist, err := biz.cartStore.FindCartItem(ctx, map[string]interface{}{"user_id": data.UserId, "food_id": data.FoodId})
		if err != nil {
			return err
		}
		if exist.UserId != 0 {
			if err := biz.cartStore.UpdateCartItem(ctx, &cartmodel.CartItemUpdate{UserId: data.UserId, FoodId: data.FoodId, Quantity: exist.Quantity + data.Quantity}); err != nil {
				return err
			}
			return nil
		}
	}

	data.RestaurantId = food.RestaurantId

	//cart of user is empty, add item to cart
	if err := biz.cartStore.CreateCartItem(ctx, data); err != nil {
		return err
	}
	return nil
}
