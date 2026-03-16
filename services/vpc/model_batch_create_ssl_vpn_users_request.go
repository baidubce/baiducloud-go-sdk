

package vpc

type BatchCreateSslVpnUsersRequest struct {
    VpnId *string `json:"-"`
    ClientToken *string `json:"-"`
    SslVpnUsers []*SslVpnUser `json:"sslVpnUsers,omitempty"`
}