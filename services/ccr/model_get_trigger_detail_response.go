package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTriggerDetailResponse struct {
	bce.BaseResponse
	CreationTime *string          `json:"creationTime,omitempty"`
	Description  *string          `json:"description,omitempty"`
	Enabled      *bool            `json:"enabled,omitempty"`
	EventTypes   []*string        `json:"eventTypes,omitempty"`
	Filters      []*TriggerFilter `json:"filters,omitempty"`
	Id           *int32           `json:"id,omitempty"`
	Name         *string          `json:"name,omitempty"`
	Targets      []*TriggerTarget `json:"targets,omitempty"`
	UpdateTime   *string          `json:"updateTime,omitempty"`
}
