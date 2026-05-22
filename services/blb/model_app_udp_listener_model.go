package blb

type AppUDPListenerModel struct {
	ListenerPort *int32  `json:"listenerPort,omitempty"`
	Scheduler    *string `json:"scheduler,omitempty"`
	Description  *string `json:"description,omitempty"`
}
