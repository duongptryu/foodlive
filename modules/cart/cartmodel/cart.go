package cartmodel

import (
	"foodlive/common"
	"time"
)

type CartItem struct {
	UserId    int               `json:"user_id" gorm:"user_id"`
	FoodId    int               `json:"food_id" gorm:"food_id"`
	Food      common.SimpleFood `json:"-" gorm:"preload:false"`
	Quantity  int               `json:"quantity" gorm:"quantity"`
	Status    bool              `json:"status" gorm:"status"`
	CreatedAt time.Time         `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" gorm:"updated_at"`
}

func (CartItem) TableName() string {
	return "carts"
}

type CartItemCreate struct {
	UserId    int        `json:"user_id" gorm:"user_id"`
	FoodId    int        `json:"food_id" gorm:"food_id"`
	Quantity  int        `json:"quantity" gorm:"quantity"`
	Status    bool       `json:"-" gorm:"status"`
	CreatedAt *time.Time `json:"-" gorm:"created_at"`
	UpdatedAt *time.Time `json:"-" gorm:"updated_at"`
}

func (CartItemCreate) TableName() string {
	return CartItem{}.TableName()
}

func (data *CartItemCreate) Validate() error {
	//if data.Quantity < 1 {
	//	return
	//}
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
	//if data.Quantity < 1 {
	//	return
	//}
	return nil
}

var (
	ErrFoodInAnotherRestaurant = common.NewCustomError(nil, "Foods need to be in the same restaurant", "ErrFoodInAnotherRestaurant")
	ErrItemAlreadyExist        = common.NewCustomError(nil, "item already exist in cart", "ErrItemAlreadyExist")
	ErrItemDoesNotExist        = common.NewCustomError(nil, "item does not exist in cart", "ErrItemDoesNotExist")
)
