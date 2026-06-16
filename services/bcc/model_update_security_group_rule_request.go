package bcc

type UpdateSecurityGroupRuleRequest struct {
	SgVersion           *int32  `json:"-"`
	SecurityGroupRuleId *string `json:"securityGroupRuleId,omitempty"`
	Remark              *string `json:"remark,omitempty"`
	PortRange           *string `json:"portRange,omitempty"`
	SourceIp            *string `json:"sourceIp,omitempty"`
	SourceGroupId       *string `json:"sourceGroupId,omitempty"`
	DestIp              *string `json:"destIp,omitempty"`
	DestGroupId         *string `json:"destGroupId,omitempty"`
	Protocol            *string `json:"protocol,omitempty"`
}
