package cfs

type ModifyTheMountTargetPermissionGroupRequest struct {
	FsId            *string `json:"-"`
	MountID         *string `json:"-"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
}
