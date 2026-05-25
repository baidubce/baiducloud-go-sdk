package vpc

type SessionConfig struct {
	TcpTimeout  *int32 `json:"tcpTimeout,omitempty"`
	UdpTimeout  *int32 `json:"udpTimeout,omitempty"`
	IcmpTimeout *int32 `json:"icmpTimeout,omitempty"`
}
