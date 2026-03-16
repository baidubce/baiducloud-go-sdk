

package vpc

type DeleteSslVpnServerRequest struct {
    VpnId *string `json:"-"`
    SslVpnServerId *string `json:"-"`
    ClientToken *string `json:"-"`
}