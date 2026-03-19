

package vpc




type Vpn struct {
	VpnId *string `json:"vpnId,omitempty"`
	VpnName *string `json:"vpnName,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	ExpiredTime *string `json:"expiredTime,omitempty"`
	PaymentTiming *string `json:"paymentTiming,omitempty"`
	Eip *string `json:"eip,omitempty"`
	MaxConnection *int32 `json:"maxConnection,omitempty"`
	BandwidthInMbps *int32 `json:"bandwidthInMbps,omitempty"`
	VpcId *string `json:"vpcId,omitempty"`
	VpnConnNum *int32 `json:"vpnConnNum,omitempty"`
	VpnConns []*VpnConn `json:"vpnConns,omitempty"`
	SslVpnServer *SslVpnServer `json:"sslVpnServer,omitempty"`
	DeleteProtect *bool `json:"deleteProtect,omitempty"`
	Tags []*TagModel `json:"tags,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
}

