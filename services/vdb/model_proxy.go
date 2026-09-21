package vdb

type Proxy struct {
	FixedIp    *string  `json:"fixedIp,omitempty"`
	FlavorInGB *float64 `json:"flavorInGB,omitempty"`
	FloatingIp *string  `json:"floatingIp,omitempty"`
	NodeShowID *string  `json:"nodeShowID,omitempty"`
	Port       *int32   `json:"port,omitempty"`
}
