package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListL3AccessRuleResponse struct {
	bce.BaseResponse
	RequestId  *string           `json:"requestId,omitempty"`
	Result     []*AccessRuleInfo `json:"result,omitempty"`
	TotalCount *int32            `json:"totalCount,omitempty"`
	PageNo     *int32            `json:"pageNo,omitempty"`
	PageSize   *int32            `json:"pageSize,omitempty"`
}
