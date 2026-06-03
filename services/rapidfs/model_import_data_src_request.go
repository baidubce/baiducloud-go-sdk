package rapidfs

type ImportDataSrcRequest struct {
	Action         *string `json:"-"`
	ClientToken    *string `json:"-"`
	InstanceId     *string `json:"instanceId,omitempty"`
	DataSrcName    *string `json:"dataSrcName,omitempty"`
	Bucket         *string `json:"bucket,omitempty"`
	OtherAccount   *bool   `json:"otherAccount,omitempty"`
	BucketAK       *string `json:"bucketAK,omitempty"`
	BucketSK       *string `json:"bucketSK,omitempty"`
	BucketStsToken *string `json:"bucketStsToken,omitempty"`
	DirPrefix      *string `json:"dirPrefix,omitempty"`
	KeepSymlink    *bool   `json:"keepSymlink,omitempty"`
	AuthGroupId    *string `json:"authGroupId,omitempty"`
	Description    *string `json:"description,omitempty"`
	QuotaMiB       *int32  `json:"quotaMiB,omitempty"`
}
