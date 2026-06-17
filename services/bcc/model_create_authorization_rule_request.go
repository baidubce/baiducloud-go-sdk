package bcc

type CreateAuthorizationRuleRequest struct {
	Action                         *string   `json:"-"`
	EnableRule                     *int32    `json:"enableRule,omitempty"`
	AuthorizeMaintenanceOperations []*string `json:"authorizeMaintenanceOperations,omitempty"`
	RuleName                       *string   `json:"ruleName,omitempty"`
	ServerEventCategory            *string   `json:"serverEventCategory,omitempty"`
}
