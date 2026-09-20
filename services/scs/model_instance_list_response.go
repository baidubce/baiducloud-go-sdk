package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type InstanceListResponse struct {
	bce.BaseResponse
	Marker      *string          `json:"marker,omitempty"`
	MaxKeys     *int32           `json:"maxKeys,omitempty"`
	IsTruncated *bool            `json:"isTruncated,omitempty"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
	Instances   []*InstanceModel `json:"instances,omitempty"`
}
