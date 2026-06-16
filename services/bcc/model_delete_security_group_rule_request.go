package bcc

type DeleteSecurityGroupRuleRequest struct {
	SecurityGroupRuleId *string `json:"-"`
	SgVersion           *int32  `json:"-"`
}
