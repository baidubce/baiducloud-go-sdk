package cprom

type ListNotificationPoliciesRequest struct {
	PageSize    *int32  `json:"-"`
	PageNo      *int32  `json:"-"`
	KeywordType *string `json:"-"`
	Keyword     *string `json:"-"`
}
