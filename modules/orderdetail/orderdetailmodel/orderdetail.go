package orderdetailmodel

import "foodlive/common"

type OrderDetail struct {
	common.SQLModel
	UserId       int                `json:"user_id" gorm:"user_id"`
	OrderId      int                `json:"order_id" gorm:"order_id"`
	RestaurantId int                `json:"restaurant_id" gorm:"restaurant_id"`
	FoodOrigin   *common.FoodOrigin `json:"food_origin" gorm:"food_origin"`
	Price        float64            `json:"price" gorm:"price"`
	Quantity     int                `json:"quantity" gorm:"quantity"`
}

func (OrderDetail) TableName() string {
	return "orders_detail"
}

type OrderDetailCreate struct {
	common.SQLModelCreate
	UserId       int                `json:"user_id" gorm:"user_id"`
	OrderId      int                `json:"order_id" gorm:"order_id"`
	RestaurantId int                `json:"restaurant_id" gorm:"restaurant_id"`
	FoodOrigin   *common.FoodOrigin `json:"food_origin" gorm:"food_origin"`
	Price        float64            `json:"price" gorm:"price"`
	Quantity     int                `json:"quantity" gorm:"quantity"`
}

func (OrderDetailCreate) TableName() string {
	return OrderDetail{}.TableName()
}

type OrderDetailUpdate struct {
	common.SQLModelUpdate
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetail{}.TableName()
}

func (data *OrderDetailUpdate) Validate() error {
	return nil
}
