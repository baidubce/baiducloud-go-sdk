package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListHaVipResponse struct {
	bce.BaseResponse
	HaVips      []*HaVip `json:"haVips,omitempty"`
	Marker      *string  `json:"marker,omitempty"`
	IsTruncated *bool    `json:"isTruncated,omitempty"`
	NextMarker  *string  `json:"nextMarker,omitempty"`
	MaxKeys     *int32   `json:"maxKeys,omitempty"`
}
