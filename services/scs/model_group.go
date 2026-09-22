package scs

type Group struct {
	SecurityGroupRemark *string `json:"securityGroupRemark,omitempty"`
	SecurityGroupName   *string `json:"securityGroupName,omitempty"`
	SecurityGroupId     *string `json:"securityGroupId,omitempty"`
	SecurityGroupUuid   *string `json:"securityGroupUuid,omitempty"`
	VpcId               *string `json:"vpcId,omitempty"`
	VpcName             *string `json:"vpcName,omitempty"`
	Outbound            []*Rule `json:"outbound,omitempty"`
}
