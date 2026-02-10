package eip

type TbspProtocolBlockingModel struct {
	Ip               *string                  `json:"ip,omitempty"`
	ProtocolPortList []*TbspProtocolPortModel `json:"protocolPortList,omitempty"`
}
