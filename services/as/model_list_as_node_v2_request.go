package as

type ListAsNodeV2Request struct {
	Groupid     *string `json:"-"`
	Keyword     *string `json:"-"`
	KeywordType *string `json:"-"`
	Order       *string `json:"-"`
	OrderBy     *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
}
