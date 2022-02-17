package orderdetailmodel

import "foodlive/common"

type OrderDetail struct {
	common.SQLModel
	UserId         int     `json:"user_id" gorm:"user_id"`
	OrderId        int     `json:"order_id" gorm:"order_id"`
	AddrUserOrigin string  `json:"addr_user_origin" gorm:"addr_user_origin"`
	Price          float64 `json:"price" gorm:"price"`
	Quantity       int     `json:"quantity" gorm:"quantity"`
	Discount       float64 `json:"discount" gorm:"discount"`
	Status         bool    `json:"bool" gorm:"bool"`
	PaymentType    string  `json:"payment_type" gorm:"payment_type"`
}

func (OrderDetail) TableName() string {
	return "orders_detail"
}

type OrderDetailCreate struct {
	common.SQLModelCreate
	UserId         int     `json:"user_id" gorm:"user_id"`
	OrderId        int     `json:"order_id" gorm:"order_id"`
	AddrUserOrigin string  `json:"addr_user_origin" gorm:"addr_user_origin"`
	Price          float64 `json:"price" gorm:"price"`
	Quantity       int     `json:"quantity" gorm:"quantity"`
	Discount       float64 `json:"discount" gorm:"discount"`
	Status         bool    `json:"-" gorm:"bool"`
	PaymentType    string  `json:"-" gorm:"payment_type"`
}

func (OrderDetailCreate) TableName() string {
	return OrderDetail{}.TableName()
}

type OrderDetailUpdate struct {
	common.SQLModelUpdate
	Status bool   `json:"-" gorm:"bool"`
	Reason string `json:"-" gorm:"reason"`
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetail{}.TableName()
}

func (data *OrderDetailUpdate) Validate() error {
	return nil
}
