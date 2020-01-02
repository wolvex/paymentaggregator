package paggr

import (
	"encoding/json"

	_ "gopkg.in/go-playground/validator.v9"
)

type Amount struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
}

type Store struct {
	ID         int64  `json:"-"`
	Code       string `json:"code,omitempty" validate:"required,max=50,min=3"`
	Name       string `json:"name,omitempty" validate:"omitempty,max=50,min=3"`
	Address    string `json:"address,omitempty" validate:"omitempty,max=50,min=3"`
	MerchantID string `json:"merchantId,omitempty" validate:"omitempty,max=50,min=3"`
	MallID     string `json:"mallId,omitempty" validate:"omitempty,max=50,min=3"`
	TerminalID string `json:"terminalId,omitempty" validate:"omitempty,max=50,min=3"`
	PublicKey  string `json:"-"`
	NotifUrl   string `json:"-"`
	SuccessUrl string `json:"successUrl,omitempty" validate:"omitempty,max=50,min=3"`
	FailedUrl  string `json:"failedUrl,omitempty" validate:"omitempty,max=50,min=3"`
}

type GoodsItem struct {
	ID          int64     `json:"-"`
	Code        string    `json:"code,omitempty" validate:"required,max=50,min=3"`
	Description string    `json:"description,omitempty" validate:"omitempty,max=255,min=1"`
	Category    string    `json:"category,omitempty" validate:"omitempty,max=50,min=3"`
	Price       *Amount   `json:"price,omitempty"`
	Qty         int       `json:"quantity,omitempty" validate:"required,numeric,max=1000000,min=1"`
	Cancellable int       `json:"-"`
	Merchant    *Merchant `json:"merchant,omitempty"`
	CheckAvail  int       `json:"-"`
	StockAvail  int       `json:"stockAvail,omitempty"`
	ExtendData  string    `json:"extendData,omitempty"`
}

type Merchant struct {
	ID        int64  `json:"-"`
	Code      string `json:"code,omitempty" validate:"required,max=50,min=3"`
	Url       string `json:"-"`
	PublicKey string `json:"-"`
	Timeout   int64  `json:"-"`
}

type Order struct {
	ID            int64  `json:"id,omitempty" validate:"omitempty,numeric,max=9999999999,min=1"`
	Title         string `json:"title,omitempty" validate:"required,max=255,min=3"`
	CustomerID    string `json:"customerId,omitempty" validate:"required,max=15,min=10"`
	CustomerName  string `json:"customerName,omitempty" validate:"omitempty,max=30,min=3"`
	CustomerEmail string `json:"customerEmail,omitempty" validate:"omitempty,email"`
	InvoiceNo     string `json:"invoiceNo,omitempty" validate:"omitempty,max=30,min=3"`
	//Timestamp     string       `json:"timestamp,omitempty" validate:"omitempty,max=50,min=10"`
	TotalPrice *Amount      `json:"totalPrice,omitempty"`
	Goods      []*GoodsItem `json:"goods,omitempty"`
	Status     int          `json:"-"`
	Remark     string       `json:"-"`
}

type Partner struct {
	ID            int64  `json:"-"`
	Code          string `json:"code,omitempty" validate:"required,max=50,min=3"`
	IssuerCode    string `json:"issuerCode,omitempty" validate:"required,max=50,min=3"`
	Hotline       string `json:"hotline,omitempty" validate:"omitempty,max=50,min=3"`
	InvoiceTmpl   int    `json:"-"`
	Url           string `json:"-"`
	PublicKey     string `json:"-"`
	AsyncPayment  int    `json:"-"`
	OrderLifetime int    `json:"-"`
}

type Account struct {
	ID            string `json:"id,omitempty" validate:"required,max=50,min=3"`
	WalletID      string `json:"walletId,omitempty" validate:"omitempty,max=50,min=3"`
	Name          string `json:"name,omitempty" validate:"omitempty,max=50,min=3"`
	Authorization string `json:"authorization,omitempty" validate:"omitempty,max=50,min=3"`
}

type Payment struct {
	ID           int64     `json:"paymentId,omitempty"`
	Method       string    `json:"method,omitempty" validate:"required,max=50,min=3"`
	Reference    string    `json:"reference,omitempty" validate:"required,max=50,min=3"`
	SequenceID   int64     `json:"sequenceId,omitempty" validate:"omitempty,numeric,max=9999999999,min=1"`
	BatchID      string    `json:"batchId,omitempty" validate:"omitempty,max=9,min=1"`
	Account      *Account  `json:"account,omitempty"`
	Partner      *Partner  `json:"partner,omitempty"`
	Resource     []*Amount `json:"resource,omitempty"`
	Reward       []*Amount `json:"reward,omitempty"`
	Balance      []*Amount `json:"balance,omitempty"`
	ApprovalCode string    `json:"approvalCode,omitempty" validate:"omitempty,max=50,min=3"`
	Token        string    `json:"token,omitempty" validate:"omitempty,max=50,min=3"`
	RedirectUrl  string    `json:"redirectUrl,omitempty" validate:"omitempty,max=50,min=3"`
	ExpiryTime   string    `json:"expiryTime,omitempty" validate:"omitempty,max=50,min=3"`
	Status       int       `json:"-"`
	Remark       string    `json:"-"`
	ForceAdvise  int       `json:"forceAdvice,omitempty"`
}

type Void struct {
	ID           int64     `json:"voidId,omitempty"`
	Method       string    `json:"method,omitempty" validate:"required,max=50,min=3"`
	Reference    string    `json:"reference,omitempty" validate:"required,max=50,min=3"`
	SequenceID   int64     `json:"sequenceId,omitempty" validate:"omitempty,numeric,max=9999999999,min=1"`
	BatchID      string    `json:"batchId,omitempty" validate:"omitempty,max=9,min=1"`
	Account      *Account  `json:"account,omitempty"`
	Partner      *Partner  `json:"partner,omitempty"`
	Resource     []*Amount `json:"resource,omitempty"`
	Reward       []*Amount `json:"reward,omitempty"`
	Balance      []*Amount `json:"balance,omitempty"`
	ApprovalCode string    `json:"approvalCode,omitempty" validate:"omitempty,max=50,min=3"`
	ReasonCode   int       `json:"reasonCode,omitempty" validate:"omitempty,numeric,max=99,min=1"`
	Status       int       `json:"-"`
	Remark       string    `json:"-"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

type RequestMessage struct {
	Store    *Store   `json:"store,omitempty"`
	Order    *Order   `json:"order,omitempty"`
	Payment  *Payment `json:"payment,omitempty"`
	Void     *Void    `json:"void,omitempty"`
	Result   *Result  `json:"result,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type ResponseMessage struct {
	Store    *Store   `json:"store,omitempty"`
	Order    *Order   `json:"order,omitempty"`
	Payment  *Payment `json:"payment,omitempty"`
	Void     *Void    `json:"void,omitempty"`
	Result   *Result  `json:"result,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type Message struct {
	Request    *RequestMessage  `json:"request,omitempty"`
	Response   *ResponseMessage `json:"response,omitempty"`
	Signature  string           `json:"signature,omitempty"`
	OriginHost string           `json:"-"`
	Version    string           `json:"-"`
	MsgID      string           `json:"-"`
	Payload    json.RawMessage  `json:"-"`
}

type Payload struct {
	Request  json.RawMessage `json:"request,omitempty"`
	Response json.RawMessage `json:"response,omitempty"`
}
