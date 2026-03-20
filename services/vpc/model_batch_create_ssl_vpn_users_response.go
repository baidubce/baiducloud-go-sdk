package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchCreateSslVpnUsersResponse struct {
	bce.BaseResponse
	SslVpnUserIds []*string `json:"sslVpnUserIds,omitempty"`
}
