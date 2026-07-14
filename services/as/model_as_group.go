package as

type AsGroup struct {
	GroupId          *string           `json:"groupId,omitempty"`
	GroupName        *string           `json:"groupName,omitempty"`
	Region           *string           `json:"region,omitempty"`
	Status           *string           `json:"status,omitempty"`
	Tags             []*TagInfo        `json:"tags,omitempty"`
	RelationTag      *bool             `json:"relationTag,omitempty"`
	VpcId            *string           `json:"vpcId,omitempty"`
	ZoneInfo         []*interface{}    `json:"zoneInfo,omitempty"`
	Config           *GroupConfig      `json:"config,omitempty"`
	Blb              []*BlbInfo        `json:"blb,omitempty"`
	NodeNum          *int32            `json:"nodeNum,omitempty"`
	RuleCount        *int32            `json:"ruleCount,omitempty"`
	CreateTime       *string           `json:"createTime,omitempty"`
	HealthCheckState *HealthCheckState `json:"healthCheckState,omitempty"`
}
