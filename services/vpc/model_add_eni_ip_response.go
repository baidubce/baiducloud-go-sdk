package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddEniIpResponse struct {
	bce.BaseResponse
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
}
