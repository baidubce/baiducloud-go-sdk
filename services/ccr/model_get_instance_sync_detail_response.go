package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceSyncDetailResponse struct {
	bce.BaseResponse
	CreationTime      *string                 `json:"creationTime,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	DestInstanceId    *string                 `json:"destInstanceId,omitempty"`
	DestProjectName   *string                 `json:"destProjectName,omitempty"`
	DestRegion        *string                 `json:"destRegion,omitempty"`
	Id                *int32                  `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Override          *bool                   `json:"override,omitempty"`
	SrcProjectName    *string                 `json:"srcProjectName,omitempty"`
	SrcRegion         *string                 `json:"srcRegion,omitempty"`
	SrcRepositoryName *string                 `json:"srcRepositoryName,omitempty"`
	SrcTagName        *string                 `json:"srcTagName,omitempty"`
	SyncType          *string                 `json:"syncType,omitempty"`
	Trigger           *ReplicationSyncTrigger `json:"trigger,omitempty"`
	UpdateTime        *string                 `json:"updateTime,omitempty"`
}
