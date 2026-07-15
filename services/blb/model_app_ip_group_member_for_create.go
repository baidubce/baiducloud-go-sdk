package blb

type AppIpGroupMemberForCreate struct {
	Ip     *string `json:"ip,omitempty"`
	Port   *int32  `json:"port,omitempty"`
	Weight *int32  `json:"weight,omitempty"`
	Desc   *string `json:"desc,omitempty"`
}
