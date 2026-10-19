package rds

type TransformationExampleRequest struct {
	InstanceId          *string      `json:"-"`
	CpuCount            *int32       `json:"cpuCount,omitempty"`
	MemoryCapacity      *int32       `json:"memoryCapacity,omitempty"`
	VolumeCapacity      *int32       `json:"volumeCapacity,omitempty"`
	AllocatedMemoryInMB *int32       `json:"allocatedMemoryInMB,omitempty"`
	NodeAmount          *int32       `json:"nodeAmount,omitempty"`
	IsDirectPay         *bool        `json:"isDirectPay,omitempty"`
	IsEnhanced          *bool        `json:"isEnhanced,omitempty"`
	EffectiveTime       *string      `json:"effectiveTime,omitempty"`
	MasterAzone         *string      `json:"masterAzone,omitempty"`
	BackupAzone         *string      `json:"backupAzone,omitempty"`
	DiskIoType          *string      `json:"diskIoType,omitempty"`
	SubnetId            *string      `json:"subnetId,omitempty"`
	EdgeSubnetId        *string      `json:"edgeSubnetId,omitempty"`
	Subnets             []*SubnetMap `json:"subnets,omitempty"`
}
