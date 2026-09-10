package cce

type InstancePage struct {
	ClusterID    *string        `json:"clusterID,omitempty"`
	KeywordType  *string        `json:"keywordType,omitempty"`
	Keyword      *string        `json:"keyword,omitempty"`
	OrderBy      *string        `json:"orderBy,omitempty"`
	Order        *string        `json:"order,omitempty"`
	PageNo       *int32         `json:"pageNo,omitempty"`
	PageSize     *int32         `json:"pageSize,omitempty"`
	TotalCount   *int32         `json:"totalCount,omitempty"`
	InstanceList []*interface{} `json:"instanceList,omitempty"`
}
