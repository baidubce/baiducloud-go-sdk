package aihc

type RetrieveTheDatasetVersionListV2Request struct {
	DatasetId  *string `json:"-"`
	PageNumber *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
