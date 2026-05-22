package vpc

type EnterpriseSecurityGroupRuleModel struct {
	Remark                        *string `json:"remark,omitempty"`
	Direction                     *string `json:"direction,omitempty"`
	Ethertype                     *string `json:"ethertype,omitempty"`
	PortRange                     *string `json:"portRange,omitempty"`
	SourcePortRange               *string `json:"sourcePortRange,omitempty"`
	Protocol                      *string `json:"protocol,omitempty"`
	SourceIp                      *string `json:"sourceIp,omitempty"`
	DestIp                        *string `json:"destIp,omitempty"`
	LocalIp                       *string `json:"localIp,omitempty"`
	RemoteIpSet                   *string `json:"remoteIpSet,omitempty"`
	RemoteIpGroup                 *string `json:"remoteIpGroup,omitempty"`
	Action                        *string `json:"action,omitempty"`
	Priority                      *int32  `json:"priority,omitempty"`
	EnterpriseSecurityGroupRuleId *string `json:"enterpriseSecurityGroupRuleId,omitempty"`
	CreatedTime                   *string `json:"createdTime,omitempty"`
	UpdatedTime                   *string `json:"updatedTime,omitempty"`
}
