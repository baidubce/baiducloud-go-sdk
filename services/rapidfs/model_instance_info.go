package rapidfs

type InstanceInfo struct {
	InstanceName           *string  `json:"instanceName,omitempty"`
	InstanceId             *string  `json:"instanceId,omitempty"`
	Description            *string  `json:"description,omitempty"`
	Status                 *string  `json:"status,omitempty"`
	Region                 *string  `json:"region,omitempty"`
	Zone                   *string  `json:"zone,omitempty"`
	CapacityMiB            *int64   `json:"capacityMiB,omitempty"`
	CapacityUsedMiB        *int64   `json:"capacityUsedMiB,omitempty"`
	CapacityUsedPercentage *float64 `json:"capacityUsedPercentage,omitempty"`
	VpcId                  *string  `json:"vpcId,omitempty"`
	SubnetId               *string  `json:"subnetId,omitempty"`
	ManagedMode            *string  `json:"managedMode,omitempty"`
	MetaSpec               *string  `json:"metaSpec,omitempty"`
	DataSpec               *string  `json:"dataSpec,omitempty"`
	RapidfsType            *string  `json:"type,omitempty"`
	ResizeStatus           *string  `json:"resizeStatus,omitempty"`
	AllocatedQuotaMiB      *int32   `json:"allocatedQuotaMiB,omitempty"`
	CceClusterId           *string  `json:"cceClusterId,omitempty"`
	AihcResourcePoolId     *string  `json:"aihcResourcePoolId,omitempty"`
	K8sControllerId        *string  `json:"k8sControllerId,omitempty"`
	Tags                   []*Tag   `json:"tags,omitempty"`
	CreateTime             *string  `json:"createTime,omitempty"`
}
