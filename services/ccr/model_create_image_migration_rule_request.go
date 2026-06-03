package ccr

type CreateImageMigrationRuleRequest struct {
	InstanceId      *string                     `json:"-"`
	Description     *string                     `json:"description,omitempty"`
	DestProjectName *string                     `json:"destProjectName,omitempty"`
	Filters         []*ReplicationFilterRequest `json:"filters,omitempty"`
	Name            *string                     `json:"name,omitempty"`
	Override        *bool                       `json:"override,omitempty"`
	SrcRegistry     *ReplicationRegistryRequest `json:"srcRegistry,omitempty"`
	Trigger         *ReplicationTriggerRequest  `json:"trigger,omitempty"`
}
