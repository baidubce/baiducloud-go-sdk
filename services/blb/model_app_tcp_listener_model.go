package blb

type AppTCPListenerModel struct {
	ListenerPort      *int32  `json:"listenerPort,omitempty"`
	Scheduler         *string `json:"scheduler,omitempty"`
	TcpSessionTimeout *int32  `json:"tcpSessionTimeout,omitempty"`
	Description       *string `json:"description,omitempty"`
}
