package vpc

type CreateVpnTunnelRequest struct {
	VpnId         *string      `json:"-"`
	ClientToken   *string      `json:"-"`
	SecretKey     *string      `json:"secretKey,omitempty"`
	LocalSubnets  []*string    `json:"localSubnets,omitempty"`
	CgwId         *string      `json:"cgwId,omitempty"`
	RemoteSubnets []*string    `json:"remoteSubnets,omitempty"`
	Description   *string      `json:"description,omitempty"`
	VpnConnName   *string      `json:"vpnConnName,omitempty"`
	IkeConfig     *IkeConfig   `json:"ikeConfig,omitempty"`
	IpsecConfig   *IpsecConfig `json:"ipsecConfig,omitempty"`
}
