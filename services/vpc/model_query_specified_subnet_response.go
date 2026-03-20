package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QuerySpecifiedSubnetResponse struct {
	bce.BaseResponse
	Subnet *SubnetDetail `json:"subnet,omitempty"`
}
