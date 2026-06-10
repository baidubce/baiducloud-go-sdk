package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetVolumeResponse struct {
	bce.BaseResponse
	Volume *VolumeModel `json:"volume,omitempty"`
}
