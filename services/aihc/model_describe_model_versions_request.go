package aihc

type DescribeModelVersionsRequest struct {
	ModelId    *string `json:"-"`
	PageNumber *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
