package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryVpcIntranetIpResponse struct {
	bce.BaseResponse
	VpcPrivateIpAddresses []*VpcPrivateIpAddress `json:"vpcPrivateIpAddresses,omitempty"`
}
