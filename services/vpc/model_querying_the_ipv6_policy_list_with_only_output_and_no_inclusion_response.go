package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionResponse struct {
	bce.BaseResponse
	EgressOnlyRules []*EgressOnlyRule `json:"egressOnlyRules,omitempty"`
	Marker          *string           `json:"marker,omitempty"`
	IsTruncated     *bool             `json:"isTruncated,omitempty"`
	NextMarker      *string           `json:"nextMarker,omitempty"`
	MaxKeys         *int32            `json:"maxKeys,omitempty"`
}
