package bci

type InstanceModel struct {
	InstanceId      *string  `json:"instanceId,omitempty"`
	InstanceName    *string  `json:"instanceName,omitempty"`
	Status          *string  `json:"status,omitempty"`
	ZoneName        *string  `json:"zoneName,omitempty"`
	CpuType         *string  `json:"cpuType,omitempty"`
	GpuType         *string  `json:"gpuType,omitempty"`
	Cpu             *float32 `json:"cpu,omitempty"`
	Memory          *float32 `json:"memory,omitempty"`
	BandwidthInMbps *int32   `json:"bandwidthInMbps,omitempty"`
	InternalIp      *string  `json:"internalIp,omitempty"`
	PublicIp        *string  `json:"publicIp,omitempty"`
	CreateTime      *string  `json:"createTime,omitempty"`
	UpdateTime      *string  `json:"updateTime,omitempty"`
	DeleteTime      *string  `json:"deleteTime,omitempty"`
	RestartPolicy   *string  `json:"restartPolicy,omitempty"`
	Tags            []*Tag   `json:"tags,omitempty"`
}
