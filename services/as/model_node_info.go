package as

type NodeInfo struct {
	Spec               *string          `json:"spec,omitempty"`
	CpuCount           *int32           `json:"cpuCount,omitempty"`
	MemoryCapacityInGB *int32           `json:"memoryCapacityInGB,omitempty"`
	SysDiskType        *string          `json:"sysDiskType,omitempty"`
	SysDiskInGB        *int32           `json:"sysDiskInGB,omitempty"`
	Billing            *BillingInfo     `json:"billing,omitempty"`
	BidModel           *string          `json:"bidModel,omitempty"`
	BidPrice           *string          `json:"bidPrice,omitempty"`
	EphemeralDisks     []*EphemeralDisk `json:"ephemeralDisks,omitempty"`
	InstanceType       *string          `json:"instanceType,omitempty"`
	GpuCard            *string          `json:"gpuCard,omitempty"`
	GpuCount           *int32           `json:"gpuCount,omitempty"`
	FpgaCard           *string          `json:"fpgaCard,omitempty"`
	FpgaCount          *int32           `json:"fpgaCount,omitempty"`
	ContainsFpga       *bool            `json:"containsFpga,omitempty"`
	KunlunCard         *string          `json:"kunlunCard,omitempty"`
	KunlunCount        *int32           `json:"kunlunCount,omitempty"`
	ImageType          *string          `json:"imageType,omitempty"`
	ImageId            *string          `json:"imageId,omitempty"`
	ImageName          *string          `json:"imageName,omitempty"`
	OsType             *string          `json:"osType,omitempty"`
	OsName             *string          `json:"osName,omitempty"`
	OsVersion          *string          `json:"osVersion,omitempty"`
	OsArch             *string          `json:"osArch,omitempty"`
	SecurityGroupId    *string          `json:"securityGroupId,omitempty"`
	AdminPass          *string          `json:"adminPass,omitempty"`
	AdminPassType      *string          `json:"adminPassType,omitempty"`
	SecurityGroupName  *string          `json:"securityGroupName,omitempty"`
	TotalCount         *int32           `json:"totalCount,omitempty"`
	AspId              *string          `json:"aspId,omitempty"`
	Cds                []*CdsInfo       `json:"cds,omitempty"`
	ZoneSubnet         *string          `json:"zoneSubnet,omitempty"`
	UserData           *string          `json:"userData,omitempty"`
	Priorities         *int32           `json:"priorities,omitempty"`
	TemplateId         *string          `json:"templateId,omitempty"`
}
