package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListInstanceSyncsResponse struct {
	bce.BaseResponse
	Total    *int32                         `json:"total,omitempty"`
	PageNo   *int32                         `json:"pageNo,omitempty"`
	PageSize *int32                         `json:"pageSize,omitempty"`
	Items    []*ReplicationSyncPolicyResult `json:"items,omitempty"`
}
