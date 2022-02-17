package historypaymentmodel

import "foodlive/common"

const (
	EntityName = "HistoryPayment"
)

type HistoryPayment struct {
	common.SQLModel
	OrderId      int    `json:"order_id" gorm:"order_id"`
	RequestId    int    `json:"request_id" gorm:"request_id"`
	OrderInfo    string `json:"order_info" gorm:"order_info"`
	OrderType    string `json:"order_type" gorm:"order_type"`
	ErrorCode    int    `json:"error_code" gorm:"error_code"`
	Message      string `json:"message" gorm:"message"`
	LocalMessage string `json:"local_message" gorm:"local_message"`
	State        string `json:"state" gorm:"state"`
	Status       bool   `json:"status" gorm:"status"`
}

func (HistoryPayment) TableName() string {
	return "orders_history"
}

type HistoryPaymentCreate struct {
	common.SQLModelCreate
	OrderId      int    `json:"order_id" gorm:"order_id"`
	RequestId    int    `json:"request_id" gorm:"request_id"`
	OrderInfo    string `json:"order_info" gorm:"order_info"`
	OrderType    string `json:"order_type" gorm:"order_type"`
	ErrorCode    int    `json:"error_code" gorm:"error_code"`
	Message      string `json:"message" gorm:"message"`
	LocalMessage string `json:"local_message" gorm:"local_message"`
	State        string `json:"state" gorm:"state"`
	Status       bool   `json:"status" gorm:"status"`
}

func (HistoryPaymentCreate) TableName() string {
	return HistoryPayment{}.TableName()
}

type HistoryPaymentUpdate struct {
	common.SQLModelUpdate
	ErrorCode    int    `json:"error_code" gorm:"error_code"`
	Message      string `json:"message" gorm:"message"`
	LocalMessage string `json:"local_message" gorm:"local_message"`
	State        string `json:"state" gorm:"state"`
	Status       bool   `json:"status" gorm:"status"`
}

func (HistoryPaymentUpdate) TableName() string {
	return HistoryPayment{}.TableName()
}