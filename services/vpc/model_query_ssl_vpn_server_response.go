

package vpc


import "github.com/baidubce/baiducloud-go-sdk/bce"



type QuerySslVpnServerResponse struct {
	bce.BaseResponse
	VpnId *string `json:"vpnId,omitempty"`
	SslVpnServerId *string `json:"sslVpnServerId,omitempty"`
	SslVpnServerName *string `json:"sslVpnServerName,omitempty"`
	InterfaceType *string `json:"interfaceType,omitempty"`
	Status *string `json:"status,omitempty"`
	LocalSubnets []*string `json:"localSubnets,omitempty"`
	RemoteSubnet *string `json:"remoteSubnet,omitempty"`
	ClientDns *string `json:"clientDns,omitempty"`
	MaxConnection *int32 `json:"maxConnection,omitempty"`
}

