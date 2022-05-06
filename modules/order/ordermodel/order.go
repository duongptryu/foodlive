package ordermodel

import (
	"foodlive/common"
	"foodlive/modules/orderdetail/orderdetailmodel"
	"foodlive/modules/ordertracking/ordertrackingmodel"
)

const (
	EntityName = "Order"
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
	UserId         int                               `json:"user_id" gorm:"user_id"`
	User           *common.SimpleUser                `json:"user" gorm:"preload:false"`
	TotalPrice     float64                           `json:"total_price" gorm:"total_price"`
	ShipperId      int                               `json:"shipper_id" gorm:"shipper_id"`
	UserAddressOri string                            `json:"user_address_ori" gorm:"user_address_ori"`
	UserAddressId  int                               `json:"user_address_id" gorm:"user_address_id"`
	UrlPayment     string                            `json:"url_payment" gorm:"url_payment"`
	Status         bool                              `json:"status" gorm:"status"`
	TypePayment    string                            `json:"type_payment" gorm:"column:type_payment"`
	TxnHash        string                            `json:"txn_hash" gorm:"column:txn_hash"`
	TotalPriceEth  string                            `json:"total_price_eth" gorm:"column:total_price_eth"`
	RestaurantId   int                               `json:"restaurant_id" gorm:"restaurant_id"`
	Restaurant     *common.SimpleRst                 `json:"restaurant" gorm:"preload:false"`
	OrderTracking  *ordertrackingmodel.OrderTracking `json:"order_tracking" gorm:"references:OrderId;foreignKey:Id;preload:false"`
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
	UserAddressId  int     `json:"user_address_id" gorm:"user_address_id"`
	UrlPayment     string  `json:"url_payment" gorm:"url_payment"`
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
	Status        *bool  `json:"-" gorm:"status"`
	TxnHash       string `json:"-" gorm:"column:txn_hash"`
	UrlPayment    string `json:"-" gorm:"url_payment"`
	TotalPriceEth string `json:"-" gorm:"column:total_price_eth"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}

type OrderGroupByUser struct {
	UserId int                `json:"user_id" gorm:"column:user_id"`
	Count  int                `json:"count" gorm:"column:count"`
	User   *common.SimpleUser `json:"user" gorm:"preload:false"`
}

type WebHookPayment struct {
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

func (data *WebHookPayment) Validate() error {
	return nil
}

type FoodQuantity struct {
	*common.SimpleFood `json:",inline"`
	Quantity           int `json:"quantity"`
}

type PreviewOrder struct {
	Foods         []FoodQuantity `json:"foods"`
	ShipFee       int            `json:"ship_fee"`
	Distance      float64        `json:"distance"`
	TotalPrice    float64        `json:"total_price"`
	TotalPriceEth string         `json:"total_price_eth"`
}

type PaymentCoinResp struct {
	OrderId int    `json:"order_id"`
	Web     string `json:"web"`
	App     string `json:"app"`
}

type OrderResponse struct {
	Order         *Order                            `json:"order"`
	OrderDetail   []orderdetailmodel.OrderDetail    `json:"order_detail"`
	OrderTracking *ordertrackingmodel.OrderTracking `json:"order_tracking"`
}

type PaymentOrderEvent struct {
	OrderId     int
	Value       string
	Hash        string
	BlockNumber uint64
}

var ErrPaymentFailed = common.NewFullErrorResponse(409, nil, "Cannot get payment, please try again!", "Cannot get payment, please try again!", "ErrPaymentFailed")
var ErrCartEmpty = common.NewCustomError(nil, "Cart is empty", "ErrCartEmpty")
var ErrOrderExpire = common.NewCustomError(nil, "Order expire", "ErrOrderExpire")
var ErrOrderAlreadyPrepare = common.NewCustomError(nil, "Order already prepare", "ErrOrderAlreadyPrepare")
var ErrMoneyTooBig = common.NewCustomError(nil, "Number of money too big, please choose another type payment", "ErrMoneyTooBig")
