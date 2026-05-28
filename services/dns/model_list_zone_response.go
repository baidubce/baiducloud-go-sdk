package dns

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListZoneResponse struct {
	bce.BaseResponse
	Marker      *string       `json:"marker,omitempty"`
	IsTruncated *bool         `json:"isTruncated,omitempty"`
	NextMarker  *string       `json:"nextMarker,omitempty"`
	MaxKeys     *int32        `json:"maxKeys,omitempty"`
	Zones       []*PublicZone `json:"zones,omitempty"`
}
