package aihc

type DescribeDatasetVersionsRequest struct {
	DatasetId  *string `json:"-"`
	PageNumber *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
