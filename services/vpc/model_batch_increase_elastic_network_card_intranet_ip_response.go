package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchIncreaseElasticNetworkCardIntranetIpResponse struct {
	bce.BaseResponse
	PrivateIpAddresses []*string `json:"privateIpAddresses,omitempty"`
}
