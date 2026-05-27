package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTgwRouteResponse struct {
	bce.BaseResponse
	TgwRtRules  []*TgwRtRule `json:"tgwRtRules,omitempty"`
	Marker      *string      `json:"marker,omitempty"`
	IsTruncated *bool        `json:"isTruncated,omitempty"`
	NextMarker  *string      `json:"nextMarker,omitempty"`
	MaxKeys     *int32       `json:"maxKeys,omitempty"`
}
