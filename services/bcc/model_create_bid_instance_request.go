package bcc

type CreateBidInstanceRequest struct {
	Spec                      *string            `json:"spec,omitempty"`
	ImageId                   *string            `json:"imageId,omitempty"`
	Billing                   *Billing           `json:"billing,omitempty"`
	BidModel                  *string            `json:"bidModel,omitempty"`
	BidPrice                  *string            `json:"bidPrice,omitempty"`
	CpuCount                  *int32             `json:"cpuCount,omitempty"`
	MemoryCapacityInGB        *int32             `json:"memoryCapacityInGB,omitempty"`
	RootDiskSizeInGb          *int32             `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType       *string            `json:"rootDiskStorageType,omitempty"`
	CreateCdsList             []*CreateCdsModel  `json:"createCdsList,omitempty"`
	EphemeralDisks            []*EphemeralDisk   `json:"ephemeralDisks,omitempty"`
	NetworkCapacityInMbps     *int32             `json:"networkCapacityInMbps,omitempty"`
	InternetChargeType        *string            `json:"internetChargeType,omitempty"`
	EipName                   *string            `json:"eipName,omitempty"`
	PurchaseCount             *int32             `json:"purchaseCount,omitempty"`
	Name                      *string            `json:"name,omitempty"`
	Hostname                  *string            `json:"hostname,omitempty"`
	AutoSeqSuffix             *bool              `json:"autoSeqSuffix,omitempty"`
	IsOpenHostnameDomain      *bool              `json:"isOpenHostnameDomain,omitempty"`
	AdminPass                 *string            `json:"adminPass,omitempty"`
	KeypairId                 *string            `json:"keypairId,omitempty"`
	UserData                  *string            `json:"userData,omitempty"`
	ZoneName                  *string            `json:"zoneName,omitempty"`
	SubnetId                  *string            `json:"subnetId,omitempty"`
	SecurityGroupId           *string            `json:"securityGroupId,omitempty"`
	EnterpriseSecurityGroupId *string            `json:"enterpriseSecurityGroupId,omitempty"`
	IsomerismCard             *string            `json:"isomerismCard,omitempty"`
	DeletionProtection        *int32             `json:"deletionProtection,omitempty"`
	RelationTag               *bool              `json:"relationTag,omitempty"`
	IsOpenIpv6                *bool              `json:"isOpenIpv6,omitempty"`
	Tags                      []*TagModel        `json:"tags,omitempty"`
	AspId                     *string            `json:"aspId,omitempty"`
	FileSystems               []*FileSystemModel `json:"fileSystems,omitempty"`
	IsEipAutoRelatedDelete    *bool              `json:"isEipAutoRelatedDelete,omitempty"`
	ResGroupId                *string            `json:"resGroupId,omitempty"`
}
