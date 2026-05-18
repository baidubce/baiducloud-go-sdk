package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListEipGroupResponse struct {
	bce.BaseResponse
	Eipgroups   []*EipGroupModel `json:"eipgroups,omitempty"`
	Marker      *string          `json:"marker,omitempty"`
	IsTruncated *bool            `json:"isTruncated,omitempty"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
	MaxKeys     *int32           `json:"maxKeys,omitempty"`
}
