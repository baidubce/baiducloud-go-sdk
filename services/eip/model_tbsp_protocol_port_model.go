package eip

type TbspProtocolPortModel struct {
	Type      *string `json:"type,omitempty"`
	PortBegin *int32  `json:"portBegin,omitempty"`
	PortEnd   *int32  `json:"portEnd,omitempty"`
}
