package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateL2PolicyResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
