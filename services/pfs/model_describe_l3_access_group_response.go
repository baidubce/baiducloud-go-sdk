package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeL3AccessGroupResponse struct {
	bce.BaseResponse
	RequestId            *string                    `json:"requestId,omitempty"`
	AccessGroupName      *string                    `json:"accessGroupName,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	InstanceCount        *int32                     `json:"instanceCount,omitempty"`
	MountCount           *int32                     `json:"mountCount,omitempty"`
	CreateTime           *string                    `json:"createTime,omitempty"`
	ZoneName             *string                    `json:"zoneName,omitempty"`
	AttachedMountTargets []*AttachedMountTargetInfo `json:"attachedMountTargets,omitempty"`
}
