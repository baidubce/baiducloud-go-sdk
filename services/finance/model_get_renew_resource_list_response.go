package finance

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRenewResourceListResponse struct {
	bce.BaseResponse
	AccountId    *string          `json:"accountId,omitempty"`
	LoginName    *string          `json:"loginName,omitempty"`
	SubAccountId *string          `json:"subAccountId,omitempty"`
	SubLoginName *string          `json:"subLoginName,omitempty"`
	OuName       *string          `json:"ouName,omitempty"`
	PageNo       *int32           `json:"pageNo,omitempty"`
	PageSize     *int32           `json:"pageSize,omitempty"`
	TotalCount   *int32           `json:"totalCount,omitempty"`
	Resources    []*RenewResource `json:"resources,omitempty"`
}
