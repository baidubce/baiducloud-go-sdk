package blb

type AppIpGroupMember struct {
	Ip       *string                      `json:"ip,omitempty"`
	Port     *int32                       `json:"port,omitempty"`
	Weight   *int32                       `json:"weight,omitempty"`
	Desc     *string                      `json:"desc,omitempty"`
	MemberId *string                      `json:"memberId,omitempty"`
	PortList []*AppIpGroupMemberPortModel `json:"portList,omitempty"`
}
