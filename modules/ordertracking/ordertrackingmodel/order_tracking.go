package ordertrackingmodel

import "foodlive/common"

type OrderTracking struct {
	common.SQLModel
	OrderId int    `json:"order_id" gorm:"order_id"`
	State   string `json:"state" gorm:"state"`
	Status  bool   `json:"status:" gorm:"status"`
}

func (OrderTracking) TableName() string {
	return "orders_tracking"
}

type OrderTrackingCreate struct {
	common.SQLModelCreate
	OrderId int    `gorm:"order_id"`
	State   string `gorm:"state"`
	Status  bool   `gorm:"status"`
}

func (OrderTrackingCreate) TableName() string {
	return OrderTracking{}.TableName()
}

type OrderTrackingUpdate struct {
	common.SQLModelUpdate
	State  string `gorm:"state"`
	Status bool   `gorm:"status"`
}

func (OrderTrackingUpdate) TableName() string {
	return OrderTracking{}.TableName()
}
