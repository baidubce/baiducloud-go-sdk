package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDnsResolverRuleListResponse struct {
	bce.BaseResponse
	Marker      *string            `json:"marker,omitempty"`
	IsTruncated *string            `json:"isTruncated,omitempty"`
	NextMarker  *string            `json:"nextMarker,omitempty"`
	MaxKeys     *string            `json:"maxKeys,omitempty"`
	Result      []*DnsResolverRule `json:"result,omitempty"`
}
