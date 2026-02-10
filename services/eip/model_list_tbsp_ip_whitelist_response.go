package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListTbspIpWhitelistResponse struct {
	bce.BaseResponse
	IpWhitelistList []*TbspIpWhitelistModel `json:"ipWhitelistList,omitempty"`
	Marker          *string                 `json:"marker,omitempty"`
	IsTruncated     *bool                   `json:"isTruncated,omitempty"`
	NextMarker      *string                 `json:"nextMarker,omitempty"`
	MaxKeys         *int32                  `json:"maxKeys,omitempty"`
}
