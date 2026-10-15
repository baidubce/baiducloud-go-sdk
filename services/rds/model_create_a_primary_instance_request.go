package rds

type CreateAPrimaryInstanceRequest struct {
	ClientToken          *string                          `json:"-"`
	Billing              *Billing                         `json:"billing,omitempty"`
	PurchaseCount        *int32                           `json:"purchaseCount,omitempty"`
	InstanceName         *string                          `json:"instanceName,omitempty"`
	Engine               *string                          `json:"engine,omitempty"`
	EngineVersion        *string                          `json:"engineVersion,omitempty"`
	Category             *string                          `json:"category,omitempty"`
	CpuCount             *int32                           `json:"cpuCount,omitempty"`
	MemoryCapacity       *int32                           `json:"memoryCapacity,omitempty"`
	VolumeCapacity       *int32                           `json:"volumeCapacity,omitempty"`
	DiskIoType           *string                          `json:"diskIoType,omitempty"`
	ZoneNames            []*string                        `json:"zoneNames,omitempty"`
	BcmGroupName         *string                          `json:"bcmGroupName,omitempty"`
	VpcId                *string                          `json:"vpcId,omitempty"`
	IsDirectPay          *bool                            `json:"isDirectPay,omitempty"`
	Subnets              []*SubnetMap                     `json:"subnets,omitempty"`
	Tags                 []*Tag                           `json:"tags,omitempty"`
	BgwGroupId           *string                          `json:"bgwGroupId,omitempty"`
	BgwGroupExclusive    *bool                            `json:"bgwGroupExclusive,omitempty"`
	CharacterSetName     *string                          `json:"characterSetName,omitempty"`
	AutoRenewTimeUnit    *string                          `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime        *int32                           `json:"autoRenewTime,omitempty"`
	LowerCaseTableNames  *int32                           `json:"lowerCaseTableNames,omitempty"`
	ParameterTemplateId  *string                          `json:"parameterTemplateId,omitempty"`
	Ovip                 *string                          `json:"ovip,omitempty"`
	EntryPort            *int32                           `json:"entryPort,omitempty"`
	ReplicationType      *string                          `json:"replicationType,omitempty"`
	ResourceGroupId      *string                          `json:"resourceGroupId,omitempty"`
	InitialDataReference *InitialDataReference            `json:"initialDataReference,omitempty"`
	Data                 []*RecoveryToSourceInstanceModel `json:"data,omitempty"`
	LeaderInstanceId     *string                          `json:"leaderInstanceId,omitempty"`
	ResourceType         *string                          `json:"resourceType,omitempty"`
}
