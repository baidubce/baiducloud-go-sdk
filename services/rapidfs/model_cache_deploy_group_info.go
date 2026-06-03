package rapidfs

type CacheDeployGroupInfo struct {
	CacheDeployGroupName   *string     `json:"cacheDeployGroupName,omitempty"`
	InstanceId             *string     `json:"instanceId,omitempty"`
	CacheDeployGroupNS     *string     `json:"cacheDeployGroupNS,omitempty"`
	Status                 *string     `json:"status,omitempty"`
	ExpectedNum            *int32      `json:"expectedNum,omitempty"`
	RunningNum             *int32      `json:"runningNum,omitempty"`
	CapacityMiB            *int64      `json:"capacityMiB,omitempty"`
	CapacityUsedMiB        *int64      `json:"capacityUsedMiB,omitempty"`
	CapacityUsedPercentage *float64    `json:"capacityUsedPercentage,omitempty"`
	DeployPath             *string     `json:"deployPath,omitempty"`
	Config                 *string     `json:"config,omitempty"`
	CreateTime             *string     `json:"createTime,omitempty"`
	ModifyTime             *string     `json:"modifyTime,omitempty"`
	DiskInfos              []*DiskInfo `json:"diskInfos,omitempty"`
}
