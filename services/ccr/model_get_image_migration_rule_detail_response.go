package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetImageMigrationRuleDetailResponse struct {
	bce.BaseResponse
	CreationTime    *string              `json:"creationTime,omitempty"`
	Deletion        *bool                `json:"deletion,omitempty"`
	Description     *string              `json:"description,omitempty"`
	DestProjectName *string              `json:"destProjectName,omitempty"`
	DestRegistry    *ReplicationRegistry `json:"destRegistry,omitempty"`
	Enabled         *bool                `json:"enabled,omitempty"`
	ExecutionTimes  *int32               `json:"executionTimes,omitempty"`
	Filters         []*ReplicationFilter `json:"filters,omitempty"`
	Id              *int32               `json:"id,omitempty"`
	Name            *string              `json:"name,omitempty"`
	Override        *bool                `json:"override,omitempty"`
	SrcRegistry     *ReplicationRegistry `json:"srcRegistry,omitempty"`
	Trigger         *ReplicationTrigger  `json:"trigger,omitempty"`
	UpdateTime      *string              `json:"updateTime,omitempty"`
}
