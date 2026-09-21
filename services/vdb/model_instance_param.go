package vdb

type InstanceParam struct {
	AvailabilityZone     *string            `json:"availabilityZone,omitempty"`
	AzInfos              []*AzInfo          `json:"azInfos,omitempty"`
	CloneDataAppBackupId *string            `json:"cloneDataAppBackupId,omitempty"`
	CloneDataAppId       *string            `json:"cloneDataAppId,omitempty"`
	Components           []*MilvusComponent `json:"components,omitempty"`
	DataNodeNum          *int32             `json:"dataNodeNum,omitempty"`
	DiskFlavor           *int32             `json:"diskFlavor,omitempty"`
	DiskType             *string            `json:"diskType,omitempty"`
	EnableEmbedding      *bool              `json:"enableEmbedding,omitempty"`
	EnableEncryption     *bool              `json:"enableEncryption,omitempty"`
	EngineVersion        *string            `json:"engineVersion,omitempty"`
	From                 *string            `json:"from,omitempty"`
	InstanceName         *string            `json:"instanceName,omitempty"`
	InstanceNum          *int32             `json:"instanceNum,omitempty"`
	InstanceType         *string            `json:"instanceType,omitempty"`
	MasterNodeSpec       *string            `json:"masterNodeSpec,omitempty"`
	MasterNum            *int32             `json:"masterNum,omitempty"`
	NodeSpec             *string            `json:"nodeSpec,omitempty"`
	NodeType             *string            `json:"nodeType,omitempty"`
	OrderId              *string            `json:"orderId,omitempty"`
	Password             *string            `json:"password,omitempty"`
	Port                 *int32             `json:"port,omitempty"`
	ProxyNodeSpec        *string            `json:"proxyNodeSpec,omitempty"`
	ProxyNum             *int32             `json:"proxyNum,omitempty"`
	ReqSource            *string            `json:"reqSource,omitempty"`
	SubnetId             *string            `json:"subnetId,omitempty"`
	SwitchEntrance       *string            `json:"switchEntrance,omitempty"`
	VpcId                *string            `json:"vpcId,omitempty"`
}
