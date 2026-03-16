

package vpc


import "github.com/baidubce/baiducloud-go-sdk/bce"



type CreateSslVpnServerResponse struct {
	bce.BaseResponse
	SslVpnServerId *string `json:"sslVpnServerId,omitempty"`
}

