package aigw

type ListAIGatewaysRequest struct {
	Keyword         *string `json:"-"`
	KeywordType     *string `json:"-"`
	Status          *string `json:"-"`
	SrcProduct      *string `json:"-"`
	TagKey          *string `json:"-"`
	TagValue        *string `json:"-"`
	ResourceGroupId *string `json:"-"`
	PageNo          *int32  `json:"-"`
	PageSize        *int32  `json:"-"`
	OrderBy         *string `json:"-"`
	Order           *string `json:"-"`
	XRegion         *string `json:"-"`
}
