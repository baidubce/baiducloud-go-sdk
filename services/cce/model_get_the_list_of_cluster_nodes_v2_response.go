package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTheListOfClusterNodesV2Response struct {
	bce.BaseResponse
	InstancePage *InstancePage `json:"instancePage,omitempty"`
	RequestID    *string       `json:"requestID,omitempty"`
}
