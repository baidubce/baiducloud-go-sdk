

package vpc

type QuerySslVpnUsersRequest struct {
    VpnId *string `json:"-"`
    Marker *string `json:"-"`
    MaxKeys *int32 `json:"-"`
    UserName *string `json:"-"`
}