package vpc

type SslVpnServer struct {
	VpnId            *string   `json:"vpnId,omitempty"`
	SslVpnServerId   *string   `json:"sslVpnServerId,omitempty"`
	SslVpnServerName *string   `json:"sslVpnServerName,omitempty"`
	InterfaceType    *string   `json:"interfaceType,omitempty"`
	Status           *string   `json:"status,omitempty"`
	LocalSubnets     []*string `json:"localSubnets,omitempty"`
	RemoteSubnet     *string   `json:"remoteSubnet,omitempty"`
	ClientDns        *string   `json:"clientDns,omitempty"`
	MaxConnection    *int32    `json:"maxConnection,omitempty"`
}
