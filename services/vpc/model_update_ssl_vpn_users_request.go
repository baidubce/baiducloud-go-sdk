

package vpc

type UpdateSslVpnUsersRequest struct {
    VpnId *string `json:"-"`
    UserId *string `json:"-"`
    ClientToken *string `json:"-"`
    ClientName *string `json:"clientName,omitempty"`
    Password *string `json:"password,omitempty"`
    Description *string `json:"description,omitempty"`
}