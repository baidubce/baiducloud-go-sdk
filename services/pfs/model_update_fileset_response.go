package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateFilesetResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
