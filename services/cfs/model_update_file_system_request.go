package cfs

type UpdateFileSystemRequest struct {
	FsId          *string `json:"-"`
	FsName        *string `json:"fsName,omitempty"`
	CapacityQuota *int32  `json:"capacityQuota,omitempty"`
}
