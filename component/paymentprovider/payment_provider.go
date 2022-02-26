package paymentprovider

import "context"

type TransactionReq struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestType string `json:"requestType"`
	NotifyURL   string `json:"notifyUrl"`
	ReturnURL   string `json:"returnUrl"`
	OrderID     string `json:"orderId"`
	Amount      string `json:"amount"`
	OrderInfo   string `json:"orderInfo"`
	RequestID   string `json:"requestId"`
	ExtraData   string `json:"extraData"`
	Signature   string `json:"signature"`
}

type TransactionResp struct {
	ErrorCode    int    `json:"status_code"`
	LocalMessage string `json:"local_msg"`
	Message      string `json:"msg"`
	OrderID      string `json:"orderId,omitempty"`
	PayUrl       string `json:"payUrl,omitempty"`
}

type OrderRequester interface {
	GetPrice() float64
	GetOrderId() int
}

type PaymentProvider interface {
	SendRequestPayment(ctx context.Context, data OrderRequester, dataExtra string) (*TransactionResp, error)
}
