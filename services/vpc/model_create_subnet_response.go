package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSubnetResponse struct {
	bce.BaseResponse
	SubnetId *string `json:"subnetId,omitempty"`
}
