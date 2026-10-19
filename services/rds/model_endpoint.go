package rds

type Endpoint struct {
	Port     *int32  `json:"port,omitempty"`
	VnetIp   *string `json:"vnetIp,omitempty"`
	InetIp   *string `json:"inetIp,omitempty"`
	Address  *string `json:"address,omitempty"`
	InetIpV6 *string `json:"inetIpV6,omitempty"`
	VnetIpV6 *string `json:"vnetIpV6,omitempty"`
}
