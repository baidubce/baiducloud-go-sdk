package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbsResponse struct {
	bce.BaseResponse
	BlbList     []*BLBModel `json:"blbList,omitempty"`
	Marker      *string     `json:"marker,omitempty"`
	IsTruncated *bool       `json:"isTruncated,omitempty"`
	NextMarker  *string     `json:"nextMarker,omitempty"`
	MaxKeys     *int32      `json:"maxKeys,omitempty"`
}
