package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListNotificationPoliciesResponse struct {
	bce.BaseResponse
	KeywordType *string       `json:"keywordType,omitempty"`
	Keyword     *string       `json:"keyword,omitempty"`
	OrderBy     *string       `json:"orderBy,omitempty"`
	Order       *string       `json:"order,omitempty"`
	PageNo      *int32        `json:"pageNo,omitempty"`
	PageSize    *int32        `json:"pageSize,omitempty"`
	TotalCount  *int32        `json:"totalCount,omitempty"`
	Items       []*NotifyRule `json:"items,omitempty"`
}
