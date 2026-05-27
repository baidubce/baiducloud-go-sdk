package vpc

type AuthorizeSecurityGroupRulesRequest struct {
	SecurityGroupId *string                 `json:"-"`
	SgVersion       *int64                  `json:"-"`
	ClientToken     *string                 `json:"-"`
	Rule            *SecurityGroupRuleModel `json:"rule,omitempty"`
}
