package cfs

type FileSystemModel struct {
	FsId            *string        `json:"fsId,omitempty"`
	FsName          *string        `json:"fsName,omitempty"`
	VpcId           *string        `json:"vpcId,omitempty"`
	CfsType         *string        `json:"type,omitempty"`
	Protocol        *string        `json:"protocol,omitempty"`
	FsUsage         *string        `json:"fsUsage,omitempty"`
	Zone            *string        `json:"zone,omitempty"`
	Status          *string        `json:"status,omitempty"`
	KMSKeyId        *string        `json:"KMSKeyId,omitempty"`
	CreateTime      *string        `json:"createTime,omitempty"`
	CapacityQuota   *int32         `json:"capacityQuota,omitempty"`
	MountTargetList []*interface{} `json:"mountTargetList,omitempty"`
	Tags            []*interface{} `json:"tags,omitempty"`
}
