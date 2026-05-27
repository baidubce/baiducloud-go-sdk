package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryRouteTableListResponse struct {
	bce.BaseResponse
	CsnRts      []*CsnRouteTable `json:"csnRts,omitempty"`
	Marker      *string          `json:"marker,omitempty"`
	IsTruncated *bool            `json:"isTruncated,omitempty"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
	MaxKeys     *int32           `json:"maxKeys,omitempty"`
}
