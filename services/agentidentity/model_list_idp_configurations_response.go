package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListIdpConfigurationsResponse struct {
	bce.BaseResponse
	TotalCount *int32                  `json:"totalCount,omitempty"`
	PageNo     *int32                  `json:"pageNo,omitempty"`
	PageSize   *int32                  `json:"pageSize,omitempty"`
	Result     []*UserPoolIdpConfigDTO `json:"result,omitempty"`
}
