package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryMountingTargetResponse struct {
	bce.BaseResponse
	MountTargetList []*MountTargetModel `json:"mountTargetList,omitempty"`
	Marker          *string             `json:"marker,omitempty"`
	IsTruncated     *bool               `json:"isTruncated,omitempty"`
	NextMarker      *string             `json:"nextMarker,omitempty"`
	MaxKeys         *int32              `json:"maxKeys,omitempty"`
}
