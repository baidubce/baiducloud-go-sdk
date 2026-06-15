package bcc

type RdmaNicTopo struct {
	RdmaIp      *string `json:"rdmaIp,omitempty"`
	SwitchName  *string `json:"switchName,omitempty"`
	SwitchPort  *string `json:"switchPort,omitempty"`
	RdmaMac     *string `json:"rdmaMac,omitempty"`
	RdmaGateway *string `json:"rdmaGateway,omitempty"`
}
