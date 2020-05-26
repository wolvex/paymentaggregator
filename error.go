package paggr

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
	0:  200,
	1:  202,
	2:  400,
	4:  400,
	5:  400,
	6:  401,
	7:  401,
	8:  400,
	9:  400,
	10: 400,
	11: 400,
	12: 400,
	13: 400,
	14: 400,
	15: 400,
	16: 400,
	17: 400,
	18: 400,
	19: 400,
	20: 400,
	21: 400,
	22: 400,
	30: 400,
	31: 400,
	32: 401,
	33: 401,
	34: 401,
	35: 400,
	36: 400,
	37: 400,
	51: 500,
	70: 500,
	71: 500,
	80: 400,
	81: 400,
	90: 500,
	91: 500,
	92: 500,
	93: 500,
	94: 400,
	98: 202,
	99: 202,
}
