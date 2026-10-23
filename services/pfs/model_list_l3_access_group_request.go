package pfs

type ListL3AccessGroupRequest struct {
	Action      *string `json:"-"`
	Keyword     *string `json:"keyword,omitempty"`
	KeywordType *string `json:"keywordType,omitempty"`
	ZoneName    *string `json:"zoneName,omitempty"`
	Order       *string `json:"order,omitempty"`
	OrderBy     *string `json:"orderBy,omitempty"`
	PageNo      *int32  `json:"pageNo,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty"`
}
