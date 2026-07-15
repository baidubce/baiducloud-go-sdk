package blb

type AppIpGroupMemberForUpdate struct {
	MemberId *string `json:"memberId,omitempty"`
	Port     *int32  `json:"port,omitempty"`
	Weight   *int32  `json:"weight,omitempty"`
	Desc     *string `json:"desc,omitempty"`
}
