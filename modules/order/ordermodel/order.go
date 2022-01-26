package ordermodel

import "foodlive/common"

type Order struct {
	common.SQLModel
	UserId     int     `json:"user_id" gorm:"user_id"`
	TotalPrice float64 `json:"total_price" gorm:"total_price"`
	ShipperId  int     `json:"shipper_id" gorm:"shipper_id"`
	Status     bool    `json:"status" gorm:"status"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderCreate struct {
	common.SQLModelCreate
	UserId     int     `json:"user_id" gorm:"user_id"`
	TotalPrice float64 `json:"total_price" gorm:"total_price"`
	ShipperId  int     `json:"shipper_id" gorm:"shipper_id"`
	Status     bool    `json:"status" gorm:"status"`
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}

func (data *OrderCreate) Validate() error {
	return nil
}

type OrderUpdate struct {
	common.SQLModelUpdate
	Status bool `json:"status" gorm:"status"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}

func (data *OrderUpdate) Validate() error {
	return nil
}
