package paymentprovider

import "context"

type TransactionReq struct {
	PartnerCode string `json:"partnerCode"`
	PartnerName string `json:"partnerName"`
	StoreId     string `json:"storeId"`
	AccessKey   string `json:"accessKey"`
	RequestType string `json:"requestType"`
	IpnUrl      string `json:"ipnUrl"`
	RedirectUrl string `json:"redirectUrl"`
	OrderID     string `json:"orderId"`
	Amount      string `json:"amount"`
	Lang        string `json:"lang"`
	AutoCapture bool   `json:"autoCapture"`
	OrderInfo   string `json:"orderInfo"`
	RequestId   string `json:"requestId"`
	ExtraData   string `json:"extraData"`
	Signature   string `json:"signature"`
}

type TransactionResp struct {
	PartnerCode  string  `json:"partnerCode"`
	OrderId      string  `json:"orderId"`
	RequestId    string  `json:"requestId"`
	Amount       float64 `json:"amount"`
	ResponseTime float64 `json:"responseTime"`
	Message      string  `json:"message"`
	ResultCode   int     `json:"resultCode"`
	PayUrl       string  `json:"payUrl"`
	Deeplink     string  `json:"deeplink"`
	QrCodeUrl    string  `json:"qrCodeUrl"`
}

type OrderRequester interface {
	GetPrice() float64
	GetOrderId() int
}

type PaymentProvider interface {
	SendRequestPayment(ctx context.Context, data OrderRequester, dataExtra string) (*TransactionResp, error)
}

type CryptoPaymentProvider interface {
	ParsePriceToEth(ctx context.Context, priceDola float64) (float64, error)
	ParseEthToVND(ctx context.Context, eth float64) (float64, error)
	CheckStatusTxn(ctx context.Context, txnHash string) (string, error)
}
