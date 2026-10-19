package rds

type InstanceDict struct {
	Port    *int32  `json:"port,omitempty"`
	VnetIp  *string `json:"vnetIp,omitempty"`
	InetIp  *string `json:"inetIp,omitempty"`
	Address *string `json:"address,omitempty"`
}
