package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTheListOfClusterNodeGroupsV2Response struct {
	bce.BaseResponse
	Page      *ListInstanceGroupPage `json:"page,omitempty"`
	RequestID *string                `json:"requestID,omitempty"`
}
