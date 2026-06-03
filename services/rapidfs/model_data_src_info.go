package rapidfs

type DataSrcInfo struct {
	DataSrcName   *string `json:"dataSrcName,omitempty"`
	DataSrcId     *string `json:"dataSrcId,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	MountTarget   *string `json:"mountTarget,omitempty"`
	Status        *string `json:"status,omitempty"`
	Description   *string `json:"description,omitempty"`
	Bucket        *string `json:"bucket,omitempty"`
	DirPrefix     *string `json:"dirPrefix,omitempty"`
	KeepSymlink   *bool   `json:"keepSymlink,omitempty"`
	AuthGroupId   *string `json:"authGroupId,omitempty"`
	AuthGroupName *string `json:"authGroupName,omitempty"`
	QuotaMiB      *int32  `json:"quotaMiB,omitempty"`
	UsedMiB       *int32  `json:"usedMiB,omitempty"`
}
