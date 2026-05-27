package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTgwListResponse struct {
	bce.BaseResponse
	Tgws        []*Tgw  `json:"tgws,omitempty"`
	Marker      *string `json:"marker,omitempty"`
	IsTruncated *bool   `json:"isTruncated,omitempty"`
	NextMarker  *string `json:"nextMarker,omitempty"`
	MaxKeys     *int32  `json:"maxKeys,omitempty"`
}
