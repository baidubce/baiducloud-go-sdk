

package vpc

type CreateSslVpnServerRequest struct {
    VpnId *string `json:"-"`
    ClientToken *string `json:"-"`
    SslVpnServerName *string `json:"sslVpnServerName,omitempty"`
    InterfaceType *string `json:"interfaceType,omitempty"`
    LocalSubnets []*string `json:"localSubnets,omitempty"`
    RemoteSubnet *string `json:"remoteSubnet,omitempty"`
    ClientDns *string `json:"clientDns,omitempty"`
}