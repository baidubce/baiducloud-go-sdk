package blb

type BlbSecurityGroupModel struct {
	SecurityGroupId    *string                      `json:"securityGroupId,omitempty"`
	SecurityGroupName  *string                      `json:"securityGroupName,omitempty"`
	SecurityGroupDesc  *string                      `json:"securityGroupDesc,omitempty"`
	VpcName            *string                      `json:"vpcName,omitempty"`
	SecurityGroupRules []*BlbSecurityGroupRuleModel `json:"securityGroupRules,omitempty"`
}
