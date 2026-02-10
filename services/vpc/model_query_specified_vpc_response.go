package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QuerySpecifiedVpcResponse struct {
	bce.BaseResponse
	Vpc *ShowVpcModel `json:"vpc,omitempty"`
}
