package cprom

type ListPodMonitorsRequest struct {
	InstanceId  *string `json:"-"`
	AgentId     *string `json:"-"`
	PageNo      *int32  `json:"-"`
	PageSize    *int32  `json:"-"`
	KeywordType *string `json:"-"`
	Keyword     *string `json:"-"`
	OrderBy     *string `json:"-"`
	Order       *string `json:"-"`
}
