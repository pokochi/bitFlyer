package bitFlyer

const (
	ME_CHILDORDERS_PATH         = "/v1/me/getchildorders"
	ME_PARENTORDERS_PATH        = "/v1/me/getparentorders"
	ME_PARENTORDER_PATH         = "/v1/me/getparentorder"
	ME_EXECUTIONS_PATH          = "/v1/me/getexecutions"
	ME_POSITIONS_PATH           = "/v1/me/getpositions"
	ME_SENDCHILDORDER_PATH      = "/v1/me/sendchildorder"
	ME_CANCELCHILDORDER_PATH    = "/v1/me/cancelchildorder"
	ME_CANCELPARENTORDER_PATH   = "/v1/me/cancelparentorder"
	ME_CANCELALLCHILDORDER_PATH = "/v1/me/cancelallchildorders"
)

type Trade struct {
	Lightning
}

func NewTrade(key string, secret string) *Trade {
	l := Lightning{key, secret}
	return &Trade{l}
}

//注文の一覧を取得
func (t *Trade) GetChildOrders(m map[string]string) (error, string) {
	return t.RequestPrivate(GET_METHOD, ME_CHILDORDERS_PATH, m)
}

//親注文の一覧を取得
func (t *Trade) GetParentOrders(m map[string]string) (error, string) {
	return t.RequestPrivate(GET_METHOD, ME_PARENTORDERS_PATH, m)
}

//親注文の詳細を取得
func (t *Trade) GetParentOrder(m map[string]string) (error, string) {
	return t.RequestPrivate(GET_METHOD, ME_PARENTORDER_PATH, m)
}

//約定の一覧を取得
func (t *Trade) GetExecutions(m map[string]string) (error, string) {
	return t.RequestPrivate(GET_METHOD, ME_EXECUTIONS_PATH, m)
}

//建玉の一覧を取得
func (t *Trade) GetPositions(m map[string]string) (error, string) {
	return t.RequestPrivate(GET_METHOD, ME_POSITIONS_PATH, m)
}

//新規注文を出す
func (t *Trade) SendChildOrder(m map[string]string) (error, string) {
	return t.RequestPrivate(POST_METHOD, ME_SENDCHILDORDER_PATH, m)
}

//注文をキャンセルする
func (t *Trade) CancelChildOrder(m map[string]string) (error, string) {
	return t.RequestPrivate(POST_METHOD, ME_CANCELCHILDORDER_PATH, m)
}

//新規の親注文を出す（特殊注文）
//TODO

//親注文をキャンセルする
func (t *Trade) CancelParentOrder(m map[string]string) (error, string) {
	return t.RequestPrivate(POST_METHOD, ME_CANCELPARENTORDER_PATH, m)
}

//すべての注文をキャンセルする
func (t *Trade) CancelAllChildOrders(m map[string]string) (error, string) {
	return t.RequestPrivate(POST_METHOD, ME_CANCELALLCHILDORDER_PATH, m)
}
