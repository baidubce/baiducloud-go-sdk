package bcc

type RevokeSecurityGroupRuleRequest struct {
	SecurityGroupId *string                 `json:"-"`
	SgVersion       *int32                  `json:"-"`
	Rule            *SecurityGroupRuleModel `json:"rule,omitempty"`
}
