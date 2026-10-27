package rapidfs

type DataSrcInfo struct {
	DataSrcName          *string                `json:"dataSrcName,omitempty"`
	DataSrcId            *string                `json:"dataSrcId,omitempty"`
	InstanceId           *string                `json:"instanceId,omitempty"`
	InstanceName         *string                `json:"instanceName,omitempty"`
	MountTarget          *string                `json:"mountTarget,omitempty"`
	Status               *string                `json:"status,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Bucket               *string                `json:"bucket,omitempty"`
	OtherAccount         *bool                  `json:"otherAccount,omitempty"`
	StorageType          *string                `json:"storageType,omitempty"`
	BucketEndpoint       *string                `json:"bucketEndpoint,omitempty"`
	DirPrefix            *string                `json:"dirPrefix,omitempty"`
	SyncMetaOnImport     *bool                  `json:"syncMetaOnImport,omitempty"`
	KeepSymlink          *bool                  `json:"keepSymlink,omitempty"`
	AuthGroupId          *string                `json:"authGroupId,omitempty"`
	AuthGroupName        *string                `json:"authGroupName,omitempty"`
	QuotaMiB             *int32                 `json:"quotaMiB,omitempty"`
	UsedMiB              *int32                 `json:"usedMiB,omitempty"`
	MountServiceBindings []*MountServiceBinding `json:"mountServiceBindings,omitempty"`
}
