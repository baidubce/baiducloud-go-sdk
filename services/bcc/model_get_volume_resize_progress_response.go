package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetVolumeResizeProgressResponse struct {
	bce.BaseResponse
	Progress *int32 `json:"progress,omitempty"`
}
