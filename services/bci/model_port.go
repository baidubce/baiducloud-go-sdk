package bci

type Port struct {
	Name     *string `json:"name,omitempty"`
	Port     *int32  `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}
