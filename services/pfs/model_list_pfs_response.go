package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListPfsResponse struct {
	bce.BaseResponse
	IsTruncated *bool            `json:"isTruncated,omitempty"`
	Marker      *string          `json:"marker,omitempty"`
	MaxKeys     *int32           `json:"maxKeys,omitempty"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
	Result      []*InstanceModel `json:"result,omitempty"`
}
