package bcc

type SecurityGroupRuleModel struct {
	Remark              *string `json:"remark,omitempty"`
	Direction           *string `json:"direction,omitempty"`
	Ethertype           *string `json:"ethertype,omitempty"`
	PortRange           *string `json:"portRange,omitempty"`
	Protocol            *string `json:"protocol,omitempty"`
	SourceGroupId       *string `json:"sourceGroupId,omitempty"`
	SourceIp            *string `json:"sourceIp,omitempty"`
	DestGroupId         *string `json:"destGroupId,omitempty"`
	DestIp              *string `json:"destIp,omitempty"`
	SecurityGroupId     *string `json:"securityGroupId,omitempty"`
	SecurityGroupRuleId *string `json:"securityGroupRuleId,omitempty"`
	CreatedTime         *string `json:"createdTime,omitempty"`
	UpdatedTime         *string `json:"updatedTime,omitempty"`
}
