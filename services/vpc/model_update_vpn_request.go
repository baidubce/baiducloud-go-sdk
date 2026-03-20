package vpc

type UpdateVpnRequest struct {
	VpnId       *string `json:"-"`
	ClientToken *string `json:"-"`
	VpnName     *string `json:"vpnName,omitempty"`
	Description *string `json:"description,omitempty"`
}
