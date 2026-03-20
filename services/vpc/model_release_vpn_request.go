package vpc

type ReleaseVpnRequest struct {
	VpnId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
