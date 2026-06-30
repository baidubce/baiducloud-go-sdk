package bci

type TCPSocketAction struct {
	Port *int32  `json:"port,omitempty"`
	Host *string `json:"host,omitempty"`
}
