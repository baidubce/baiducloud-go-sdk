package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListTriggersResponse struct {
	bce.BaseResponse
	Total    *int32           `json:"total,omitempty"`
	PageNo   *int32           `json:"pageNo,omitempty"`
	PageSize *int32           `json:"pageSize,omitempty"`
	Policies []*TriggerPolicy `json:"policies,omitempty"`
}
