package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryAclRulesResponse struct {
	bce.BaseResponse
	Marker      *string    `json:"marker,omitempty"`
	IsTruncated *bool      `json:"isTruncated,omitempty"`
	NextMarker  *string    `json:"nextMarker,omitempty"`
	MaxKeys     *int32     `json:"maxKeys,omitempty"`
	AclRules    []*AclRule `json:"aclRules,omitempty"`
}
