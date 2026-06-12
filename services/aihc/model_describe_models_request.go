package aihc

type DescribeModelsRequest struct {
	Keyword    *string `json:"-"`
	PageNumber *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
