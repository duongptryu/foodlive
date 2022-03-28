package cartmodel

import (
	"foodlive/common"
	"time"
)

type CartItem struct {
	UserId       int                `json:"-" gorm:"user_id"`
	FoodId       int                `json:"food_id" gorm:"food_id"`
	RestaurantId int                `json:"restaurant_id" gorm:"restaurant_id"`
	Food         *common.SimpleFood `json:"food" gorm:"preload:false"`
	Quantity     int                `json:"quantity" gorm:"quantity"`
	Status       bool               `json:"status" gorm:"status"`
	CreatedAt    time.Time          `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" gorm:"updated_at"`
}

func (CartItem) TableName() string {
	return "carts"
}

type CartItemCreate struct {
	UserId       int        `json:"-" gorm:"user_id"`
	FoodId       int        `json:"food_id" gorm:"food_id"`
	RestaurantId int        `json:"-" gorm:"restaurant_id"`
	Quantity     int        `json:"quantity" gorm:"quantity"`
	Status       bool       `json:"-" gorm:"status"`
	CreatedAt    *time.Time `json:"-" gorm:"created_at"`
	UpdatedAt    *time.Time `json:"-" gorm:"updated_at"`
}

func (CartItemCreate) TableName() string {
	return CartItem{}.TableName()
}

func (data *CartItemCreate) Validate() error {
	if data.Quantity < 1 {
		return ErrInvalidQuantity
	}
	data.Status = true
	return nil
}

type CartItemUpdate struct {
	UserId    int        `json:"-" gorm:"user_id"`
	FoodId    int        `json:"food_id" gorm:"food_id"`
	Quantity  int        `json:"quantity" gorm:"quantity"`
	UpdatedAt *time.Time `json:"-" gorm:"updated_at"`
}

func (CartItemUpdate) TableName() string {
	return CartItem{}.TableName()
}

func (data *CartItemUpdate) Validate() error {
	if data.Quantity < 1 {
		return ErrInvalidQuantity
	}
	return nil
}

var (
	ErrFoodInAnotherRestaurant = common.NewCustomError(nil, "Foods need to be in the same restaurant", "ErrFoodInAnotherRestaurant")
	ErrItemAlreadyExist        = common.NewCustomError(nil, "item already exist in cart", "ErrItemAlreadyExist")
	ErrItemDoesNotExist        = common.NewCustomError(nil, "item does not exist in cart", "ErrItemDoesNotExist")
	ErrInvalidQuantity         = common.NewCustomError(nil, "quantity must greater than 0", "ErrInvalidQuantity")
)
