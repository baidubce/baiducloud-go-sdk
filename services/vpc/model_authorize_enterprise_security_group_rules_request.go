package vpc

type AuthorizeEnterpriseSecurityGroupRulesRequest struct {
	EnterpriseSecurityGroupId *string                             `json:"-"`
	ClientToken               *string                             `json:"-"`
	Rules                     []*EnterpriseSecurityGroupRuleModel `json:"rules,omitempty"`
}
