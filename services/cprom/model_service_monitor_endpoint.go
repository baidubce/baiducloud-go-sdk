package cprom

type ServiceMonitorEndpoint struct {
	Interval *string `json:"interval,omitempty"`
	Path     *string `json:"path,omitempty"`
	Port     *string `json:"port,omitempty"`
}
