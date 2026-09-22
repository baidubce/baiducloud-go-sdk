package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDeploymentSetListResponse struct {
	bce.BaseResponse
	Result      []*Deploy `json:"result,omitempty"`
	Marker      *string   `json:"marker,omitempty"`
	IsTruncated *bool     `json:"isTruncated,omitempty"`
	NextMarker  *string   `json:"nextMarker,omitempty"`
	MaxKeys     *int32    `json:"maxKeys,omitempty"`
}
