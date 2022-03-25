package ordertrackingmodel

import "foodlive/common"

const (
	StateWaitingShipper = "waiting_for_shipper"
	StatePreparing      = "preparing"
	StateOnTheWay       = "on_the_way"
	StateDelivered      = "delivered"
	StateCancel         = "cancel"
	StatePaymentFail    = "payment_fail"
	StateWaitingPayment = "waiting_for_payment"
)

type OrderTracking struct {
	common.SQLModel
	OrderId int    `json:"order_id" gorm:"order_id"`
	State   string `json:"state" gorm:"state"`
}

func (OrderTracking) TableName() string {
	return "order_trackings"
}

type OrderTrackingCreate struct {
	common.SQLModelCreate
	OrderId int    `gorm:"order_id"`
	State   string `gorm:"state"`
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
