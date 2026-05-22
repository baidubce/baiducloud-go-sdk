package vpc

type AuthorizedEnterpriseSecurityGroupRulesRequest struct {
	EnterpriseSecurityGroupId *string                             `json:"-"`
	ClientToken               *string                             `json:"-"`
	Rules                     []*EnterpriseSecurityGroupRuleModel `json:"rules,omitempty"`
}
