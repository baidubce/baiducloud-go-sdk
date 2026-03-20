package vpc

type SearchVpnTunnelRequest struct {
	VpnId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
