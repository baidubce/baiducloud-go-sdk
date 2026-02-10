package eip

type AddTbspProtocolBlockingRequest struct {
	Id               *string                  `json:"-"`
	ClientToken      *string                  `json:"-"`
	Ip               *string                  `json:"ip,omitempty"`
	ProtocolPortList []*TbspProtocolPortModel `json:"protocolPortList,omitempty"`
}
