package historypaymentmodel

import "foodlive/common"

const (
	EntityName = "HistoryPayment"
)

type HistoryPayment struct {
	common.SQLModel
	Amount       int     `json:"amount"`
	OrderInfo    string  `json:"orderInfo"`
	OrderType    string  `json:"orderType"`
	TransID      float64 `json:"transId"`
	Message      string  `json:"message"`
	PartnerCode  string  `json:"partnerCode"`
	RequestID    string  `json:"requestId"`
	OrderID      string  `json:"orderId"`
	ResponseTime float64 `json:"responseTime"`
	ExtraData    string  `json:"extraData"`
	Signature    string  `json:"signature"`
	PayType      string  `json:"payType"`
	ResultCode   int     `json:"resultCode"`
	AccessKey    string  `json:"accessKey"`
	LocalMessage string  `json:"localMessage"`
}

func (HistoryPayment) TableName() string {
	return "order_history"
}

type HistoryPaymentCreate struct {
	common.SQLModelCreate
	Amount       int     `json:"amount"`
	OrderInfo    string  `json:"orderInfo"`
	OrderType    string  `json:"orderType"`
	TransID      float64 `json:"transId"`
	Message      string  `json:"message"`
	RequestID    string  `json:"requestId"`
	OrderID      string  `json:"orderId"`
	ResponseTime float64 `json:"responseTime"`
	ExtraData    string  `json:"extraData"`
	PayType      string  `json:"payType"`
	ResultCode   int     `json:"resultCode"`
	LocalMessage string  `json:"localMessage"`
}

func (HistoryPaymentCreate) TableName() string {
	return HistoryPayment{}.TableName()
}
