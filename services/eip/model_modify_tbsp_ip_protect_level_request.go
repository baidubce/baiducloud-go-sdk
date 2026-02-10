package eip

type ModifyTbspIpProtectLevelRequest struct {
	Id           *string `json:"-"`
	ClientToken  *string `json:"-"`
	Ip           *string `json:"ip,omitempty"`
	ProtectLevel *int32  `json:"protectLevel,omitempty"`
}
