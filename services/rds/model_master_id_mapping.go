package rds

type MasterIdMapping struct {
	AppId                *string  `json:"appId,omitempty"`
	AppIdShort           *string  `json:"appIdShort,omitempty"`
	Status               *string  `json:"status,omitempty"`
	AllocatedMemoryInMB  *int32   `json:"allocatedMemoryInMB,omitempty"`
	CpuCount             *int32   `json:"cpuCount,omitempty"`
	AllocatedStorageInGB *int32   `json:"allocatedStorageInGB,omitempty"`
	UsedStorageInGB      *float64 `json:"usedStorageInGB,omitempty"`
}
