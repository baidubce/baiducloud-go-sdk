package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetNodeGroupDetailsV2Response struct {
	bce.BaseResponse
	RequestID     *string        `json:"requestID,omitempty"`
	InstanceGroup *InstanceGroup `json:"instanceGroup,omitempty"`
}
