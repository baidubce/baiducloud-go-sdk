package as

type ListAsGroupV2Request struct {
	Keyword        *string `json:"-"`
	KeywordType    *string `json:"-"`
	SubKeywordType *string `json:"-"`
	Order          *string `json:"-"`
	OrderBy        *string `json:"-"`
	PageNo         *int32  `json:"-"`
	PageSize       *int32  `json:"-"`
}
