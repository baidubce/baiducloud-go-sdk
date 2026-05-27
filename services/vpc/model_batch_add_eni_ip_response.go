package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchAddEniIpResponse struct {
	bce.BaseResponse
	PrivateIpAddresses []*string `json:"privateIpAddresses,omitempty"`
}
