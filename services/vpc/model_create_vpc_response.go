package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateVpcResponse struct {
	bce.BaseResponse
	VpcId *string `json:"vpcId,omitempty"`
}
