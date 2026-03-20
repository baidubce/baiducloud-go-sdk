package vpc

type UnbindEipRequest struct {
	VpnId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
