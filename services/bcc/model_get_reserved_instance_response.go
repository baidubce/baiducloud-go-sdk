package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetReservedInstanceResponse struct {
	bce.BaseResponse
	TotalCount        *int32                  `json:"totalCount,omitempty"`
	Marker            *string                 `json:"marker,omitempty"`
	MaxKeys           *int32                  `json:"maxKeys,omitempty"`
	NextMarker        *string                 `json:"nextMarker,omitempty"`
	IsTruncated       *bool                   `json:"isTruncated,omitempty"`
	ReservedInstances []*ReservedInstanceInfo `json:"reservedInstances,omitempty"`
}
