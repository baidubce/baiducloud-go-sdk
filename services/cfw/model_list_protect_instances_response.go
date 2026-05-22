package cfw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListProtectInstancesResponse struct {
	bce.BaseResponse
	Marker      *string     `json:"marker,omitempty"`
	IsTruncated *bool       `json:"isTruncated,omitempty"`
	NextMarker  *string     `json:"nextMarker,omitempty"`
	MaxKeys     *int32      `json:"maxKeys,omitempty"`
	Instances   []*Instance `json:"instances,omitempty"`
}
