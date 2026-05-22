package vpc

type DeleteEnterpriseSecurityGroupRulesRequest struct {
	EnterpriseSecurityGroupRuleId *string `json:"-"`
	ClientToken                   *string `json:"-"`
}
