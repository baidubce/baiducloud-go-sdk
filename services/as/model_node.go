package as

type Node struct {
	InstanceId         *string    `json:"instanceId,omitempty"`
	InstanceUuid       *string    `json:"instanceUuid,omitempty"`
	InstanceName       *string    `json:"instanceName,omitempty"`
	FloatingIp         *string    `json:"floatingIp,omitempty"`
	InternalIp         *string    `json:"internalIp,omitempty"`
	Status             *string    `json:"status,omitempty"`
	Payment            *string    `json:"payment,omitempty"`
	CpuCount           *int32     `json:"cpuCount,omitempty"`
	MemoryCapacityInGB *int32     `json:"memoryCapacityInGB,omitempty"`
	InstanceType       *string    `json:"instanceType,omitempty"`
	SysDiskInGB        *int32     `json:"sysDiskInGB,omitempty"`
	CreateTime         *string    `json:"createTime,omitempty"`
	Eip                *AsEip     `json:"eip,omitempty"`
	SubnetType         *string    `json:"subnetType,omitempty"`
	IsProtected        *bool      `json:"isProtected,omitempty"`
	NodeType           *string    `json:"nodeType,omitempty"`
	Tags               []*TagInfo `json:"tags,omitempty"`
	GroupId            *string    `json:"groupId,omitempty"`
	IsManaged          *bool      `json:"isManaged,omitempty"`
	InternalSpec       *string    `json:"internalSpec,omitempty"`
	LogicalZone        *string    `json:"logicalZone,omitempty"`
}
