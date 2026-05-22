package vpc

type DeleteRegularSecurityGroupRulesV2Request struct {
	SecurityGroupRuleId *string `json:"-"`
	ClientToken         *string `json:"-"`
	SgVersion           *int64  `json:"-"`
}
