package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateVolumeResponse struct {
	bce.BaseResponse
	OrderId     *string   `json:"orderId,omitempty"`
	VolumeIds   []*string `json:"volumeIds,omitempty"`
	WarningList []*string `json:"warningList,omitempty"`
}
