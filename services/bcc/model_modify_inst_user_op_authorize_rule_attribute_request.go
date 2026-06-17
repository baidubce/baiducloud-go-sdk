package bcc

type ModifyInstUserOpAuthorizeRuleAttributeRequest struct {
	Action                         *string   `json:"-"`
	EnableRule                     *int32    `json:"enableRule,omitempty"`
	AuthorizeMaintenanceOperations []*string `json:"authorizeMaintenanceOperations,omitempty"`
	RuleName                       *string   `json:"ruleName,omitempty"`
	RuleId                         *string   `json:"ruleId,omitempty"`
}
