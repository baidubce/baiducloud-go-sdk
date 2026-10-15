package rds

type CreateReadOnlyInstanceRequest struct {
	ClientToken           *string      `json:"-"`
	Billing               *Billing     `json:"billing,omitempty"`
	SourceInstanceId      *string      `json:"sourceInstanceId,omitempty"`
	InstanceName          *string      `json:"instanceName,omitempty"`
	CpuCount              *int32       `json:"cpuCount,omitempty"`
	MemoryCapacity        *int32       `json:"memoryCapacity,omitempty"`
	VolumeCapacity        *int32       `json:"volumeCapacity,omitempty"`
	DiskIoType            *string      `json:"diskIoType,omitempty"`
	ZoneNames             []*string    `json:"zoneNames,omitempty"`
	VpcId                 *string      `json:"vpcId,omitempty"`
	BcmGroupName          *string      `json:"bcmGroupName,omitempty"`
	IsDirectPay           *bool        `json:"isDirectPay,omitempty"`
	Subnets               []*SubnetMap `json:"subnets,omitempty"`
	Ovip                  *string      `json:"ovip,omitempty"`
	EntryPort             *int32       `json:"entryPort,omitempty"`
	ResourceGroupId       *string      `json:"resourceGroupId,omitempty"`
	ReplicaType           *string      `json:"replicaType,omitempty"`
	IsInheritMasterAuthip *bool        `json:"isInheritMasterAuthip,omitempty"`
	PurchaseCount         *int32       `json:"purchaseCount,omitempty"`
	ResourceType          *string      `json:"resourceType,omitempty"`
	AutoRenewTimeUnit     *string      `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime         *int32       `json:"autoRenewTime,omitempty"`
}
