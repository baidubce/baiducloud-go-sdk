package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDnsResolverListResponse struct {
	bce.BaseResponse
	Marker      *string        `json:"marker,omitempty"`
	IsTruncated *string        `json:"isTruncated,omitempty"`
	NextMarker  *string        `json:"nextMarker,omitempty"`
	MaxKeys     *string        `json:"maxKeys,omitempty"`
	Result      []*DnsResolver `json:"result,omitempty"`
}
