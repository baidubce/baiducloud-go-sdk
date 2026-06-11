package aihc

type RetrieveTheDatasetListV2Request struct {
	Keyword          *string `json:"-"`
	StorageType      *string `json:"-"`
	StorageInstances *string `json:"-"`
	ImportFormat     *string `json:"-"`
	PageNumber       *int32  `json:"-"`
	PageSize         *int32  `json:"-"`
}
