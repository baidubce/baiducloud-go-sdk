package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListInstanceSyncRecordsResponse struct {
	bce.BaseResponse
	Total    *int32             `json:"total,omitempty"`
	PageSize *int32             `json:"pageSize,omitempty"`
	Items    []*ExecutionResult `json:"items,omitempty"`
}
