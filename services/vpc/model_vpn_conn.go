

package vpc




type VpnConn struct {
	VpnId *string `json:"vpnId,omitempty"`
	VpnConnId *string `json:"vpnConnId,omitempty"`
	VpnConnName *string `json:"vpnConnName,omitempty"`
	LocalIp *string `json:"localIp,omitempty"`
	SecretKey *string `json:"secretKey,omitempty"`
	LocalSubnets []*string `json:"localSubnets,omitempty"`
	RemoteIp *string `json:"remoteIp,omitempty"`
	RemoteSubnets []*string `json:"remoteSubnets,omitempty"`
	CgwId *string `json:"cgwId,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
	HealthStatus *string `json:"healthStatus,omitempty"`
	IkeConfig *IkeConfig `json:"ikeConfig,omitempty"`
	IpsecConfig *IpsecConfig `json:"ipsecConfig,omitempty"`
}

