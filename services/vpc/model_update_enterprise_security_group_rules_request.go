package vpc

type UpdateEnterpriseSecurityGroupRulesRequest struct {
	EnterpriseSecurityGroupRuleId *string `json:"-"`
	ClientToken                   *string `json:"-"`
	Remark                        *string `json:"remark,omitempty"`
	PortRange                     *string `json:"portRange,omitempty"`
	SourcePortRange               *string `json:"sourcePortRange,omitempty"`
	SourceIp                      *string `json:"sourceIp,omitempty"`
	DestIp                        *string `json:"destIp,omitempty"`
	LocalIp                       *string `json:"localIp,omitempty"`
	RemoteIpSet                   *string `json:"remoteIpSet,omitempty"`
	RemoteIpGroup                 *string `json:"remoteIpGroup,omitempty"`
	Action                        *string `json:"action,omitempty"`
	Priority                      *int32  `json:"priority,omitempty"`
	Protocol                      *string `json:"protocol,omitempty"`
}
