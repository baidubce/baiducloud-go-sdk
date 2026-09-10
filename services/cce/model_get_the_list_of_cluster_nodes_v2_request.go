package cce

type GetTheListOfClusterNodesV2Request struct {
	ClusterID   *string `json:"-"`
	KeywordType *string `json:"-"`
	Keyword     *string `json:"-"`
	OrderBy     *string `json:"-"`
	Order       *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
}
