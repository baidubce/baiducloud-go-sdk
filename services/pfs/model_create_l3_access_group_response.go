package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateL3AccessGroupResponse struct {
	bce.BaseResponse
	RequestId       *string `json:"requestId,omitempty"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
}
