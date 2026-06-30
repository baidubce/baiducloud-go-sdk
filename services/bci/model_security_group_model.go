package bci

type SecurityGroupModel struct {
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	VpcId           *string `json:"vpcId,omitempty"`
}
