package rapidfs

type ModifyDataSrcRequest struct {
	ClientToken          *string                `json:"-"`
	DataSrcId            *string                `json:"dataSrcId,omitempty"`
	InstanceId           *string                `json:"instanceId,omitempty"`
	BucketAK             *string                `json:"bucketAK,omitempty"`
	BucketSK             *string                `json:"bucketSK,omitempty"`
	KeepSymlink          *bool                  `json:"keepSymlink,omitempty"`
	AuthGroupId          *string                `json:"authGroupId,omitempty"`
	MountServiceBindings []*MountServiceBinding `json:"mountServiceBindings,omitempty"`
	Description          *string                `json:"description,omitempty"`
	QuotaMiB             *int32                 `json:"quotaMiB,omitempty"`
}
