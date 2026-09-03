package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListAIGatewaysResponse struct {
	bce.BaseResponse
	OrderBy    *string              `json:"orderBy,omitempty"`
	Order      *string              `json:"order,omitempty"`
	PageNo     *int32               `json:"pageNo,omitempty"`
	PageSize   *int32               `json:"pageSize,omitempty"`
	TotalCount *int64               `json:"totalCount,omitempty"`
	Result     []*AIGatewayListItem `json:"result,omitempty"`
}
