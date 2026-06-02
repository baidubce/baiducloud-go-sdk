package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteFilesetResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
