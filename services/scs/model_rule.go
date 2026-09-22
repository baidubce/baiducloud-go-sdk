package scs

type Rule struct {
	Id                  *string `json:"id,omitempty"`
	SecurityGroupRuleId *string `json:"securityGroupRuleId,omitempty"`
	SecurityGroupId     *string `json:"securityGroupId,omitempty"`
	SecurityGroupUuid   *string `json:"securityGroupUuid,omitempty"`
	Direction           *string `json:"direction,omitempty"`
	Ethertype           *string `json:"ethertype,omitempty"`
	Protocol            *string `json:"protocol,omitempty"`
	PortRange           *string `json:"portRange,omitempty"`
	RemoteGroupId       *string `json:"remoteGroupId,omitempty"`
	RemoteGroupName     *string `json:"remoteGroupName,omitempty"`
	RemoteIP            *string `json:"remoteIP,omitempty"`
	Name                *string `json:"name,omitempty"`
}
