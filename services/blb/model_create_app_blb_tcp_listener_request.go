package blb

type CreateAppBlbTcpListenerRequest struct {
	BlbId             *string `json:"-"`
	ClientToken       *string `json:"-"`
	ListenerPort      *int32  `json:"listenerPort,omitempty"`
	Scheduler         *string `json:"scheduler,omitempty"`
	TcpSessionTimeout *int32  `json:"tcpSessionTimeout,omitempty"`
	Description       *string `json:"description,omitempty"`
}
