package ccr

type ListInstancesRequest struct {
	PageNo       *int32  `json:"-"`
	PageSize     *int32  `json:"-"`
	KeywordType  *string `json:"-"`
	Keyword      *string `json:"-"`
	Acrossregion *string `json:"-"`
}
