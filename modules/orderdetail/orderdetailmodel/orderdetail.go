package orderdetailmodel

import "foodlive/common"

type OrderDetail struct {
	common.SQLModel
	OrderId    int                `json:"order_id" gorm:"order_id"`
	FoodOrigin *common.FoodOrigin `json:"food_origin" gorm:"food_origin"`
	Price      float64            `json:"price" gorm:"price"`
	Quantity   int                `json:"quantity" gorm:"quantity"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

type OrderDetailCreate struct {
	common.SQLModelCreate
	OrderId    int                `json:"order_id" gorm:"order_id"`
	FoodOrigin *common.FoodOrigin `json:"food_origin" gorm:"food_origin"`
	Price      float64            `json:"price" gorm:"price"`
	Quantity   int                `json:"quantity" gorm:"quantity"`
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
