package vpc

type RevokeRegularSecurityGroupRulesV2Request struct {
	SecurityGroupId *string                 `json:"-"`
	ClientToken     *string                 `json:"-"`
	SgVersion       *int64                  `json:"-"`
	Rule            *SecurityGroupRuleModel `json:"rule,omitempty"`
}
