package blb

type AppPolicy struct {
	Id                 *string    `json:"id,omitempty"`
	Desc               *string    `json:"desc,omitempty"`
	AppServerGroupId   *string    `json:"appServerGroupId,omitempty"`
	AppServerGroupName *string    `json:"appServerGroupName,omitempty"`
	AppIpGroupId       *string    `json:"appIpGroupId,omitempty"`
	AppIpGroupName     *string    `json:"appIpGroupName,omitempty"`
	FrontendPort       *int32     `json:"frontendPort,omitempty"`
	BlbType            *string    `json:"type,omitempty"`
	BackendPort        *int32     `json:"backendPort,omitempty"`
	PortType           *string    `json:"portType,omitempty"`
	Priority           *int32     `json:"priority,omitempty"`
	RuleList           []*AppRule `json:"ruleList,omitempty"`
	GroupType          *string    `json:"groupType,omitempty"`
}
