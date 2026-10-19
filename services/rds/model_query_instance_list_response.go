package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryInstanceListResponse struct {
	bce.BaseResponse
	Marker      *string     `json:"marker,omitempty"`
	MaxKeys     *int32      `json:"maxKeys,omitempty"`
	IsTruncated *bool       `json:"isTruncated,omitempty"`
	Instances   []*Instance `json:"instances,omitempty"`
	NextMarker  *string     `json:"nextMarker,omitempty"`
}
