package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListBaseDdosResponse struct {
	bce.BaseResponse
	DdosList    []*DdosModel `json:"ddosList,omitempty"`
	Marker      *string      `json:"marker,omitempty"`
	IsTruncated *bool        `json:"isTruncated,omitempty"`
	NextMarker  *string      `json:"nextMarker,omitempty"`
	MaxKeys     *int32       `json:"maxKeys,omitempty"`
}
