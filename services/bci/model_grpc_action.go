package bci

type GRPCAction struct {
	Port    *int32  `json:"port,omitempty"`
	Service *string `json:"service,omitempty"`
}
