package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListAlertEventsResponse struct {
	bce.BaseResponse
	TotalCount *int32   `json:"totalCount,omitempty"`
	PageNo     *int32   `json:"pageNo,omitempty"`
	PageSize   *int32   `json:"pageSize,omitempty"`
	Items      []*Event `json:"items,omitempty"`
	OrderBy    *string  `json:"orderBy,omitempty"`
	Order      *string  `json:"order,omitempty"`
}
