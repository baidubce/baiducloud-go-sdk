package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateImageResponse struct {
	bce.BaseResponse
	ImageId        *string   `json:"imageId,omitempty"`
	CdsSnapshotIds []*string `json:"cdsSnapshotIds,omitempty"`
}
