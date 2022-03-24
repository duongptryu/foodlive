package ordermodel

import "foodlive/common"

const (
	EntityName = "Order"
)

const (
	PaymentStatus     = "PAYMENT_STATUS"
	PrepareStatus     = "PREPARE_STATUS"
	DeliveryStatus    = "DELIVERY_STATUS"
	CompleteStatus    = "COMPLETE_STATUS"
	PaymentFailStatus = "PAYMENT_FAIL_STATUS"
)

const (
	TypeMomo   = "MOMO"
	TypeCrypto = "CRYPTO"
)

type Checkout struct {
	UserAddrId int `json:"user_addr_id"`
}

type Order struct {
	common.SQLModel
	UserId         int     `json:"user_id" gorm:"user_id"`
	TotalPrice     float64 `json:"total_price" gorm:"total_price"`
	ShipperId      int     `json:"shipper_id" gorm:"shipper_id"`
	UserAddressOri string  `json:"user_address_ori" gorm:"user_address_ori"`
	Status         bool    `json:"status" gorm:"status"`
	TypePayment    string  `json:"type_payment" gorm:"column:type_payment"`
	TxnHash        string  `json:"txn_hash" gorm:"column:txn_hash"`
	TotalPriceEth  float64 `json:"total_price_eth" gorm:"column:total_price_eth"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderCreate struct {
	common.SQLModelCreate
	UserId         int     `json:"-" gorm:"user_id"`
	RestaurantId   int     `json:"-" gorm:"restaurant_id"`
	TotalPrice     float64 `json:"total_price" gorm:"total_price"`
	ShipperId      int     `json:"-" gorm:"shipper_id"`
	UserAddressOri string  `json:"user_address_ori" gorm:"user_address_ori"`
	Status         bool    `json:"-" gorm:"status"`
	TypePayment    string  `json:"-" gorm:"column:type_payment"`
	TotalPriceEth  float64 `json:"-" gorm:"column:total_price_eth"`
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}

func (data *OrderCreate) Validate() error {
	return nil
}

func (data *OrderCreate) GetOrderId() int {
	return data.Id
}

func (data *OrderCreate) GetPrice() float64 {
	return data.TotalPrice
}

type OrderUpdate struct {
	common.SQLModelUpdate
	Status  *bool  `json:"-" gorm:"status"`
	TxnHash string `json:"txn_hash" gorm:"column:txn_hash"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}

type WebHookPayment struct {
	PartnerCode  string `json:"partnerCode"`
	RequestID    string `json:"requestId"`
	Amount       string `json:"amount"`
	OrderID      string `json:"orderId"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	ExtraData    string `json:"extraData"`
	Signature    string `json:"signature"`
	PayType      string `json:"payType"`
	ErrorCode    string `json:"errorCode"`
	AccessKey    string `json:"accessKey"`
	OrderType    string `json:"orderType"`
	OrderInfo    string `json:"orderInfo"`
	TransID      string `json:"transId"`
	LocalMessage string `json:"localMessage"`
}

func (data *WebHookPayment) Validate() error {
	return nil
}

type FoodQuantity struct {
	*common.SimpleFood `json:",inline"`
	Quantity           int `json:"quantity"`
}

type PreviewOrder struct {
	Foods         []FoodQuantity `json:"foods"`
	ShipFee       float64        `json:"ship_fee"`
	TotalPrice    float64        `json:"total_price"`
	TotalPriceEth float64        `json:"total_price_eth"`
}

type PaymentCoinResp struct {
	Web string `json:"web"`
	App string `json:"app"`
}

var ErrPaymentFailed = common.NewFullErrorResponse(409, nil, "Cannot get payment, please try again!", "Cannot get payment, please try again!", "ErrPaymentFailed")
