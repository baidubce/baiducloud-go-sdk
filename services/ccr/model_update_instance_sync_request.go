package ccr

type UpdateInstanceSyncRequest struct {
	InstanceId      *string                        `json:"-"`
	PolicyId        *string                        `json:"-"`
	Description     *string                        `json:"description,omitempty"`
	DestInstanceId  *string                        `json:"destInstanceId,omitempty"`
	DestProjectName *string                        `json:"destProjectName,omitempty"`
	Name            *string                        `json:"name,omitempty"`
	Override        *bool                          `json:"override,omitempty"`
	SrcProjectName  *string                        `json:"srcProjectName,omitempty"`
	SrcRepository   *string                        `json:"srcRepository,omitempty"`
	SrcTag          *string                        `json:"srcTag,omitempty"`
	SyncType        *string                        `json:"syncType,omitempty"`
	Trigger         *ReplicationSyncTriggerRequest `json:"trigger,omitempty"`
}
