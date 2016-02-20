package bitFlyer

const (
	ME_BALANCE_PATH    = "/v1/me/getbalance"
	ME_COLLATERAL_PATH = "/v1/me/getcollateral"
)

type Assets struct {
	Lightning
}

func NewAssets(key string, secret string) *Assets {
	l := Lightning{key, secret}
	return &Assets{l}
}

//資産残高を取得
func (a *Assets) GetBalance() (error, string) {
	return a.RequestPrivate(GET_METHOD, ME_BALANCE_PATH, nil)
}

//証拠金の状態を取得
func (a *Assets) GetCollateral() (error, string) {
	return a.RequestPrivate(GET_METHOD, ME_COLLATERAL_PATH, nil)
}
