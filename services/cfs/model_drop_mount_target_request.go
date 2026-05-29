package cfs

type DropMountTargetRequest struct {
	FsId    *string `json:"-"`
	MountId *string `json:"-"`
}
