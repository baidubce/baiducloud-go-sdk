package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryCsnListResponse struct {
	bce.BaseResponse
	Csns        []*Csn  `json:"csns,omitempty"`
	Marker      *string `json:"marker,omitempty"`
	IsTruncated *bool   `json:"isTruncated,omitempty"`
	NextMarker  *string `json:"nextMarker,omitempty"`
	MaxKeys     *int32  `json:"maxKeys,omitempty"`
}
