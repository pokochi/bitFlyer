package bitFlyer

const (
	ME_PERMISSIONS_PATH = "/v1/me/getpermissions"
)

type Permissions struct {
	Lightning
}

func NewPermissions(key string, secret string) *Permissions {
	l := Lightning{key, secret}
	return &Permissions{l}
}

//API キーの権限を取得
func (p *Permissions) GetPermissions() (error, string) {
	return p.RequestPrivate(GET_METHOD, ME_PERMISSIONS_PATH, nil)
}
