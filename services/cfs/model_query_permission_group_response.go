package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryPermissionGroupResponse struct {
	bce.BaseResponse
	AccessGroups []*AccessGroupModel `json:"accessGroups,omitempty"`
	Marker       *string             `json:"marker,omitempty"`
	IsTruncated  *bool               `json:"isTruncated,omitempty"`
	NextMarker   *string             `json:"nextMarker,omitempty"`
	MaxKeys      *int32              `json:"maxKeys,omitempty"`
}
