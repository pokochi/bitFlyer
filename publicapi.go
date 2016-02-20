package bitFlyer

const (
	HEALTH_PATH     = "/v1/gethealth"
	BOARD_PATH      = "/v1/getboard"
	TICKER_PATH     = "/v1/getticker"
	EXECUTIONS_PATH = "/v1/getexecutions"
)

type PublicApi struct {
	Lightning
}

//板情報
func (p *PublicApi) GetBoard(m map[string]string) (error, string) {
	return p.RequestPublic(GET_METHOD, BOARD_PATH, m)
}

//Ticker
func (p *PublicApi) GetTicker(m map[string]string) (error, string) {
	return p.RequestPublic(GET_METHOD, TICKER_PATH, m)
}

//約定履歴
func (p *PublicApi) GetExecutions(m map[string]string) (error, string) {
	return p.RequestPublic(GET_METHOD, EXECUTIONS_PATH, m)
}

//取引所の状態
func (p *PublicApi) GetHealth() (error, string) {
	return p.RequestPublic(GET_METHOD, HEALTH_PATH, nil)
}
