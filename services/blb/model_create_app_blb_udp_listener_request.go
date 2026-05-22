package blb

type CreateAppBlbUdpListenerRequest struct {
	BlbId             *string `json:"-"`
	ClientToken       *string `json:"-"`
	ListenerPort      *int32  `json:"listenerPort,omitempty"`
	Scheduler         *string `json:"scheduler,omitempty"`
	UdpSessionTimeout *int32  `json:"udpSessionTimeout,omitempty"`
	Description       *string `json:"description,omitempty"`
}
