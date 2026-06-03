package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeInstanceResponse struct {
	bce.BaseResponse
	InstanceInfo *InstanceInfo `json:"instanceInfo,omitempty"`
}
