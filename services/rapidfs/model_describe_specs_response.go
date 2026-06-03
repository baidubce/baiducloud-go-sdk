package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeSpecsResponse struct {
	bce.BaseResponse
	InstanceSpecInfos []*InstanceSpecInfo `json:"instanceSpecInfos,omitempty"`
}
