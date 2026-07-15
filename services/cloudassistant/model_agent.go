package cloudassistant

type Agent struct {
	Host    *Host   `json:"host,omitempty"`
	State   *string `json:"state,omitempty"`
	Version *string `json:"version,omitempty"`
}
