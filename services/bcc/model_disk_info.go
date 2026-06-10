package bcc

type DiskInfo struct {
	StorageType *string `json:"storageType,omitempty"`
	MaxDiskSize *int32  `json:"maxDiskSize,omitempty"`
	MinDiskSize *int32  `json:"minDiskSize,omitempty"`
}
