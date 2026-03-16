

package vpc

type UpdateSslVpnServerRequest struct {
    VpnId *string `json:"-"`
    SslVpnServerId *string `json:"-"`
    ClientToken *string `json:"-"`
    SslVpnServerName *string `json:"sslVpnServerName,omitempty"`
    LocalSubnets []*string `json:"localSubnets,omitempty"`
    RemoteSubnet *string `json:"remoteSubnet,omitempty"`
    ClientDns *string `json:"clientDns,omitempty"`
}