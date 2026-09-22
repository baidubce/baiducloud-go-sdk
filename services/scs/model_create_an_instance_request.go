package scs

type CreateAnInstanceRequest struct {
	ClientToken       *string           `json:"-"`
	Billing           *Billing          `json:"billing,omitempty"`
	InstanceName      *string           `json:"instanceName,omitempty"`
	NodeType          *string           `json:"nodeType,omitempty"`
	Port              *int32            `json:"port,omitempty"`
	Engine            *int32            `json:"engine,omitempty"`
	EngineVersion     *string           `json:"engineVersion,omitempty"`
	StoreType         *int32            `json:"storeType,omitempty"`
	EnableReadOnly    *int32            `json:"enableReadOnly,omitempty"`
	PurchaseCount     *int32            `json:"purchaseCount,omitempty"`
	ShardNum          *int32            `json:"shardNum,omitempty"`
	ProxyNum          *int32            `json:"proxyNum,omitempty"`
	ClusterType       *string           `json:"clusterType,omitempty"`
	DiskFlavor        *int32            `json:"diskFlavor,omitempty"`
	DiskType          *string           `json:"diskType,omitempty"`
	VpcId             *string           `json:"vpcId,omitempty"`
	ReplicationInfo   []*ReplicationMap `json:"replicationInfo,omitempty"`
	AutoRenewTimeUnit *string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     *int32            `json:"autoRenewTime,omitempty"`
	BgwGroupId        *string           `json:"bgwGroupId,omitempty"`
	ClientAuth        *string           `json:"clientAuth,omitempty"`
	Tags              []*Tag            `json:"tags,omitempty"`
	ConfTpl           *string           `json:"confTpl,omitempty"`
	ResourceGroupId   *string           `json:"resourceGroupId,omitempty"`
	AutoBackupConfig  *string           `json:"autoBackupConfig,omitempty"`
	DeployIdList      []*string         `json:"deployIdList,omitempty"`
}
