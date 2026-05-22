package blb

type CreateAppPolicy struct {
	AppServerGroupId *string          `json:"appServerGroupId,omitempty"`
	AppIpGroupId     *string          `json:"appIpGroupId,omitempty"`
	BackendPort      *int32           `json:"backendPort,omitempty"`
	PortType         *string          `json:"portType,omitempty"`
	Priority         *int32           `json:"priority,omitempty"`
	RuleList         []*CreateAppRule `json:"ruleList,omitempty"`
	Desc             *string          `json:"desc,omitempty"`
}
