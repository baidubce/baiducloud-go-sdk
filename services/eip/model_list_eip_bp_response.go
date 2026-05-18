package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListEipBpResponse struct {
	bce.BaseResponse
	BpList      []*BandwidthPackage `json:"bpList,omitempty"`
	Marker      *string             `json:"marker,omitempty"`
	IsTruncated *bool               `json:"isTruncated,omitempty"`
	NextMarker  *string             `json:"nextMarker,omitempty"`
	MaxKeys     *int32              `json:"maxKeys,omitempty"`
}
