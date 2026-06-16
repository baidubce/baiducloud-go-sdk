package bcc

type CreateSecurityGroupRequest struct {
	Name  *string                   `json:"name,omitempty"`
	Desc  *string                   `json:"desc,omitempty"`
	VpcId *string                   `json:"vpcId,omitempty"`
	Tags  []*Tag                    `json:"tags,omitempty"`
	Rules []*SecurityGroupRuleModel `json:"rules,omitempty"`
}
