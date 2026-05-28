package blb

type BlbEnterpriseSecurityGroupRuleModel struct {
	Remark                        *string `json:"remark,omitempty"`
	Direction                     *string `json:"direction,omitempty"`
	Ethertype                     *string `json:"ethertype,omitempty"`
	PortRange                     *string `json:"portRange,omitempty"`
	Protocol                      *string `json:"protocol,omitempty"`
	SourceIp                      *string `json:"sourceIp,omitempty"`
	DestIp                        *string `json:"destIp,omitempty"`
	Action                        *string `json:"action,omitempty"`
	Priority                      *int32  `json:"priority,omitempty"`
	EnterpriseSecurityGroupRuleId *string `json:"enterpriseSecurityGroupRuleId,omitempty"`
}
