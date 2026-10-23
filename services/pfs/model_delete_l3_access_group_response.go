package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteL3AccessGroupResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
