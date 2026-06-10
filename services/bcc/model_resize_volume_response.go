package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ResizeVolumeResponse struct {
	bce.BaseResponse
	WarningList []*string `json:"warningList,omitempty"`
}
