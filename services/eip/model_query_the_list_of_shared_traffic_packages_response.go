package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheListOfSharedTrafficPackagesResponse struct {
	bce.BaseResponse
	PackageList []*Package `json:"packageList,omitempty"`
	Marker      *string    `json:"marker,omitempty"`
	IsTruncated *bool      `json:"isTruncated,omitempty"`
	NextMarker  *string    `json:"nextMarker,omitempty"`
	MaxKeys     *int32     `json:"maxKeys,omitempty"`
}
