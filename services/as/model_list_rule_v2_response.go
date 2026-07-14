package as

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRuleV2Response struct {
	bce.BaseResponse
	OrderBy    *string `json:"orderBy,omitempty"`
	Order      *string `json:"order,omitempty"`
	PageNo     *int32  `json:"pageNo,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty"`
	Result     []*Rule `json:"result,omitempty"`
}
