package paggr

import "net/http"

const (
	ERR_SUCCESS              int = 0
	ERR_IN_PROGRESS          int = 1
	ERR_PARAM_MISSING        int = 2
	ERR_PARAM_ILLEGAL        int = 4
	ERR_INVALID_FORMAT       int = 5
	ERR_UNAUTHORIZED         int = 6
	ERR_INVALID_SIGNATURE    int = 7
	ERR_STORE_NOT_FOUND      int = 8
	ERR_PARTNER_NOT_FOUND    int = 9
	ERR_GOODS_NOT_FOUND      int = 10
	ERR_GOODS_NOT_AVAIL      int = 11
	ERR_UNSUPPORTED_CURRENCY int = 12
	ERR_TOTAL_PRICE_INVALID  int = 13
	ERR_INVALID_PRICE        int = 14
	ERR_MULTI_PARTNER        int = 15
	ERR_MULTI_MERCHANT       int = 16
	ERR_CUSTOMER_NOT_FOUND   int = 17
	ERR_ACCOUNT_NOT_FOUND    int = 18
	ERR_PAYMENT_DECLINED     int = 19
	ERR_PAYMENT_FAILED       int = 20
	ERR_PAYMENT_DUPLICATE    int = 21
	ERR_PAYMENT_IN_PROGRESS  int = 22
	ERR_PAYMENT_EXPIRED      int = 23
	ERR_TRX_EXPIRED          int = 30
	ERR_TRX_INVALID          int = 31
	ERR_TRX_UNAUTHORIZED     int = 32
	ERR_TRX_EXCEED_LIMIT     int = 33
	ERR_TRX_BELOW_LIMIT      int = 34
	ERR_TRX_REVERSED         int = 35
	ERR_TRX_UNRESOLVED       int = 36
	ERR_TRX_DUPLICATE        int = 37
	ERR_PAYMENT_TIMEOUT      int = 70
	ERR_FULFILLMENT_TIMEOUT  int = 71
	ERR_DELIVERY_CANCELLED   int = 80
	ERR_TERMINAL_SUSPENDED   int = 81
	ERR_TIMEOUT              int = 90
	ERR_DATABASE             int = 91
	ERR_SYSTEM_ERROR         int = 92
	ERR_OTHERS               int = 93
	ERR_VOID_CANCELLED       int = 94
	ERR_PROCESSING           int = 98
	ERR_IN_QUEUE             int = 99
)

const (
	VOID_REASON_PAYMENT_TIMEOUT    int = 1
	VOID_REASON_FULFILLMENT_FAILED int = 2
	VOID_REASON_RESOLVED_AS_FAILED int = 3
)

var HTTP_STATUS_MAP = map[int]int{
	0:  http.StatusOK,                  //ERR_SUCCESS
	1:  http.StatusOK,                  //ERR_IN_PROGRESS
	2:  http.StatusBadRequest,          //ERR_PARAM_MISSING
	4:  http.StatusBadRequest,          //ERR_PARAM_ILLEGAL
	5:  http.StatusBadRequest,          //ERR_INVALID_FORMAT
	6:  http.StatusUnauthorized,        //ERR_UNAUTHORIZED
	7:  http.StatusUnauthorized,        //ERR_INVALID_SIGNATURE
	8:  http.StatusBadRequest,          //ERR_STORE_NOT_FOUND
	9:  http.StatusBadRequest,          //ERR_PARTNER_NOT_FOUND
	10: http.StatusBadRequest,          //ERR_GOODS_NOT_FOUND
	11: http.StatusBadRequest,          //ERR_GOODS_NOT_AVAIL
	12: http.StatusBadRequest,          //ERR_UNSUPPORTED_CURRENCY
	13: http.StatusBadRequest,          //ERR_TOTAL_PRICE_INVALID
	14: http.StatusBadRequest,          //ERR_INVALID_PRICE
	15: http.StatusBadRequest,          //ERR_MULTI_PARTNER
	16: http.StatusBadRequest,          //ERR_MULTI_MERCHANT
	17: http.StatusBadRequest,          //ERR_CUSTOMER_NOT_FOUND
	18: http.StatusBadRequest,          //ERR_ACCOUNT_NOT_FOUND
	19: http.StatusBadRequest,          //ERR_PAYMENT_DECLINED
	20: http.StatusInternalServerError, //ERR_PAYMENT_FAILED
	21: http.StatusBadRequest,          //ERR_PAYMENT_DUPLICATE
	22: http.StatusBadRequest,          //ERR_PAYMENT_IN_PROGRESS
	30: http.StatusBadRequest,          //ERR_TRX_EXPIRED
	31: http.StatusBadRequest,          //ERR_TRX_INVALID
	32: http.StatusBadRequest,          //ERR_TRX_UNAUTHORIZED
	33: http.StatusBadRequest,          //ERR_TRX_EXCEED_LIMIT
	34: http.StatusBadRequest,          //ERR_TRX_BELOW_LIMIT
	35: http.StatusBadRequest,          //ERR_TRX_REVERSED
	36: http.StatusBadRequest,          //ERR_TRX_UNRESOLVED
	37: http.StatusBadRequest,          //ERR_TRX_DUPLICATE
	70: http.StatusInternalServerError, //ERR_PAYMENT_TIMEOUT
	71: http.StatusInternalServerError, //ERR_FULFILLMENT_TIMEOUT
	80: http.StatusInternalServerError, //ERR_DELIVERY_CANCELLED
	81: http.StatusBadRequest,          //ERR_TERMINAL_SUSPENDED
	90: http.StatusInternalServerError, //ERR_TIMEOUT
	91: http.StatusInternalServerError, //ERR_DATABASE
	92: http.StatusInternalServerError, //ERR_SYSTEM_ERROR
	93: http.StatusInternalServerError, //ERR_OTHERS
	94: http.StatusBadRequest,          //ERR_VOID_CANCELLED
	98: http.StatusOK,                  //ERR_PROCESSING
	99: http.StatusOK,                  //ERR_IN_QUEUE
}
