package blb

type BlbEnterpriseSecurityGroupModel struct {
	EnterpriseSecurityGroupId    *string                                `json:"enterpriseSecurityGroupId,omitempty"`
	EnterpriseSecurityGroupName  *string                                `json:"enterpriseSecurityGroupName,omitempty"`
	EnterpriseSecurityGroupDesc  *string                                `json:"enterpriseSecurityGroupDesc,omitempty"`
	EnterpriseSecurityGroupRules []*BlbEnterpriseSecurityGroupRuleModel `json:"enterpriseSecurityGroupRules,omitempty"`
}
