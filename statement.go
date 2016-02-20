package bitFlyer

const (
	ME_ADDRESSES_PATH   = "/v1/me/getaddresses"
	ME_COININS_PATH     = "/v1/me/getcoinins"
	ME_COINOUTS_PATH    = "/v1/me/getcoinouts"
	ME_DEPOSITS_PATH    = "/v1/me/getdeposits"
	ME_WITHDRAWALS_PATH = "/v1/me/getwithdrawals"
)

type Statement struct {
	Lightning
}

func NewStatement(key string, secret string) *Statement {
	l := Lightning{key, secret}
	return &Statement{l}
}

//預入用ビットコインアドレスの取得
func (s *Statement) GetAddresses() (error, string) {
	return s.RequestPrivate(GET_METHOD, ME_ADDRESSES_PATH, nil)
}

//ビットコイン預入履歴
func (s *Statement) GetCoinins() (error, string) {
	return s.RequestPrivate(GET_METHOD, ME_COININS_PATH, nil)
}

//ビットコイン送付履歴
func (s *Statement) GetCoinouts() (error, string) {
	return s.RequestPrivate(GET_METHOD, ME_COINOUTS_PATH, nil)
}

//デポジット入金履歴
func (s *Statement) GetDeposits(m map[string]string) (error, string) {
	return s.RequestPrivate(GET_METHOD, ME_DEPOSITS_PATH, m)
}

//デポジット解約履歴
func (s *Statement) GetWithdrawals(m map[string]string) (error, string) {
	return s.RequestPrivate(GET_METHOD, ME_WITHDRAWALS_PATH, m)
}
