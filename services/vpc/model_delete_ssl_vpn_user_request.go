package vpc

type DeleteSslVpnUserRequest struct {
	VpnId       *string `json:"-"`
	UserId      *string `json:"-"`
	ClientToken *string `json:"-"`
}
