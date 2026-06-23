package cprom

type ListAlertsRequest struct {
	InstanceId  *string `json:"-"`
	PageSize    *int32  `json:"-"`
	PageNo      *int32  `json:"-"`
	KeywordType *string `json:"-"`
	Keyword     *string `json:"-"`
}
