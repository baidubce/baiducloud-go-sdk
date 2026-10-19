package rds

type ReplicaRequest struct {
	InstanceId     *string `json:"instanceId,omitempty"`
	CpuCount       *int32  `json:"cpuCount,omitempty"`
	MemoryCapacity *int32  `json:"memoryCapacity,omitempty"`
	VolumeCapacity *int32  `json:"volumeCapacity,omitempty"`
	DiskIoType     *string `json:"diskIoType,omitempty"`
}
