package aihc

type GetModelListV2Request struct {
	Keyword    *string `json:"-"`
	PageNumber *int32  `json:"-"`
	PageSize   *int32  `json:"-"`
}
