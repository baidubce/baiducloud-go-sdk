package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListPodMonitorsResponse struct {
	bce.BaseResponse
	OrderBy     *string           `json:"orderBy,omitempty"`
	Order       *string           `json:"order,omitempty"`
	KeywordType *string           `json:"keywordType,omitempty"`
	Keyword     *string           `json:"keyword,omitempty"`
	PageNo      *int32            `json:"pageNo,omitempty"`
	PageSize    *int32            `json:"pageSize,omitempty"`
	TotalCount  *int32            `json:"totalCount,omitempty"`
	Status      *string           `json:"status,omitempty"`
	Items       []*PodMonitorItem `json:"items,omitempty"`
}
