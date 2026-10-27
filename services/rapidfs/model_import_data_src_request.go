package rapidfs

type ImportDataSrcRequest struct {
	ClientToken          *string                `json:"-"`
	InstanceId           *string                `json:"instanceId,omitempty"`
	DataSrcName          *string                `json:"dataSrcName,omitempty"`
	Bucket               *string                `json:"bucket,omitempty"`
	OtherAccount         *bool                  `json:"otherAccount,omitempty"`
	StorageType          *string                `json:"storageType,omitempty"`
	BucketEndpoint       *string                `json:"bucketEndpoint,omitempty"`
	BucketAK             *string                `json:"bucketAK,omitempty"`
	BucketSK             *string                `json:"bucketSK,omitempty"`
	BucketStsToken       *string                `json:"bucketStsToken,omitempty"`
	DirPrefix            *string                `json:"dirPrefix,omitempty"`
	SyncMetaOnImport     *bool                  `json:"syncMetaOnImport,omitempty"`
	KeepSymlink          *bool                  `json:"keepSymlink,omitempty"`
	AuthGroupId          *string                `json:"authGroupId,omitempty"`
	MountServiceBindings []*MountServiceBinding `json:"mountServiceBindings,omitempty"`
	Description          *string                `json:"description,omitempty"`
	QuotaMiB             *int32                 `json:"quotaMiB,omitempty"`
}
