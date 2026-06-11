package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetImageResponse struct {
	bce.BaseResponse
	Image *ImageModel `json:"image,omitempty"`
}
