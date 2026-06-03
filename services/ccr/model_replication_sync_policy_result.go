package ccr

type ReplicationSyncPolicyResult struct {
	Id                *int32                  `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	SyncType          *string                 `json:"syncType,omitempty"`
	SrcProjectName    *string                 `json:"srcProjectName,omitempty"`
	SrcRepositoryName *string                 `json:"srcRepositoryName,omitempty"`
	SrcTagName        *string                 `json:"srcTagName,omitempty"`
	SrcRegion         *string                 `json:"srcRegion,omitempty"`
	CreationTime      *string                 `json:"creationTime,omitempty"`
	UpdateTime        *string                 `json:"updateTime,omitempty"`
	DestInstanceId    *string                 `json:"destInstanceId,omitempty"`
	DestProjectName   *string                 `json:"destProjectName,omitempty"`
	DestRegion        *string                 `json:"destRegion,omitempty"`
	Trigger           *ReplicationSyncTrigger `json:"trigger,omitempty"`
	Override          *bool                   `json:"override,omitempty"`
}
