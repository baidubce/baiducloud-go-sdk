package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeInstancesResponse struct {
	bce.BaseResponse
	InstanceInfos []*InstanceInfo `json:"instanceInfos,omitempty"`
	Marker        *string         `json:"marker,omitempty"`
	IsTruncated   *bool           `json:"isTruncated,omitempty"`
	NextMarker    *string         `json:"nextMarker,omitempty"`
	MaxKeys       *int32          `json:"maxKeys,omitempty"`
}
