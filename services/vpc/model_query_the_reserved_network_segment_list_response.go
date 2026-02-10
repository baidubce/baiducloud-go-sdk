package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheReservedNetworkSegmentListResponse struct {
	bce.BaseResponse
	IpReserves  []*IpReserve `json:"ipReserves,omitempty"`
	Marker      *string      `json:"marker,omitempty"`
	IsTruncated *bool        `json:"isTruncated,omitempty"`
	NextMarker  *string      `json:"nextMarker,omitempty"`
	MaxKeys     *int32       `json:"maxKeys,omitempty"`
}
