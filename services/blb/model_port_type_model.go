package blb

type PortTypeModel struct {
	Port    *int32  `json:"port,omitempty"`
	BlbType *string `json:"type,omitempty"`
}
