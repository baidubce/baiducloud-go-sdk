package vpc

type DeleteSecurityGroupRulesRequest struct {
	SecurityGroupRuleId *string `json:"-"`
	ClientToken         *string `json:"-"`
	SgVersion           *int64  `json:"-"`
}
