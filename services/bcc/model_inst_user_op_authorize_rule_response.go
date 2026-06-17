package bcc

type InstUserOpAuthorizeRuleResponse struct {
	RuleId                         *string     `json:"ruleId,omitempty"`
	RuleName                       *string     `json:"ruleName,omitempty"`
	ServerEventCategory            *string     `json:"serverEventCategory,omitempty"`
	EffectiveScope                 *string     `json:"effectiveScope,omitempty"`
	Status                         *string     `json:"status,omitempty"`
	Tags                           []*TagModel `json:"tags,omitempty"`
	AuthorizeMaintenanceOperations []*string   `json:"authorizeMaintenanceOperations,omitempty"`
	CreateTime                     *string     `json:"createTime,omitempty"`
}
