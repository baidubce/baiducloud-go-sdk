package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateFilesetResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	FilesetId *string `json:"filesetId,omitempty"`
}
