package rds

type ObtainThePriceDifferenceOfPrepaidTransformerRequest struct {
	InstanceId           *string `json:"instanceId,omitempty"`
	CpuCount             *int32  `json:"cpuCount,omitempty"`
	AllocatedMemoryInGB  *int32  `json:"allocatedMemoryInGB,omitempty"`
	AllocatedStorageInGB *int32  `json:"allocatedStorageInGB,omitempty"`
	DiskIoType           *string `json:"diskIoType,omitempty"`
	Category             *string `json:"category,omitempty"`
}
