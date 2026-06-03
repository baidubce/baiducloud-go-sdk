package rapidfs

type CacheNodeInfo struct {
	CacheNodeId            *string            `json:"cacheNodeId,omitempty"`
	InstanceId             *string            `json:"instanceId,omitempty"`
	RapidfsType            *string            `json:"type,omitempty"`
	Ip                     *string            `json:"ip,omitempty"`
	Status                 *string            `json:"status,omitempty"`
	ConnectionStatus       *string            `json:"connectionStatus,omitempty"`
	Description            *string            `json:"description,omitempty"`
	CreateTime             *string            `json:"createTime,omitempty"`
	ReportTime             *string            `json:"reportTime,omitempty"`
	CapacityMiB            *int64             `json:"capacityMiB,omitempty"`
	CapacityUsedMiB        *int64             `json:"capacityUsedMiB,omitempty"`
	CapacityUsedPercentage *float64           `json:"capacityUsedPercentage,omitempty"`
	Config                 *string            `json:"config,omitempty"`
	DiskInfos              []*DiskInfo        `json:"diskInfos,omitempty"`
	DeployPath             *string            `json:"deployPath,omitempty"`
	BccInfo                *BCCCacheNodeInfo  `json:"bccInfo,omitempty"`
	IdcInfo                *IDCCacheNodeInfo  `json:"idcInfo,omitempty"`
	CceInfo                *CCECacheNodeInfo  `json:"cceInfo,omitempty"`
	K8sInfo                *K8SCacheNodeInfo  `json:"k8sInfo,omitempty"`
	AihcInfo               *AIHCCacheNodeInfo `json:"aihcInfo,omitempty"`
}
