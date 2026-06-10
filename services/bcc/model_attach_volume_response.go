package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AttachVolumeResponse struct {
	bce.BaseResponse
	VolumeAttachment *VolumeAttachmentModel `json:"volumeAttachment,omitempty"`
	WarningList      []*string              `json:"warningList,omitempty"`
}
