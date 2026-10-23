package pfs

type AccessGroupInfo struct {
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	Description     *string `json:"description,omitempty"`
	RuleCount       *int32  `json:"ruleCount,omitempty"`
	InstanceCount   *int32  `json:"instanceCount,omitempty"`
	MountCount      *int32  `json:"mountCount,omitempty"`
	CreateTime      *string `json:"createTime,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
}
