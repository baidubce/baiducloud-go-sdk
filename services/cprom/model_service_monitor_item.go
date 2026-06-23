package cprom

type ServiceMonitorItem struct {
	ServiceMonitorName *string     `json:"serviceMonitorName,omitempty"`
	Namespace          *string     `json:"namespace,omitempty"`
	Enable             *string     `json:"enable,omitempty"`
	CreateTime         *string     `json:"createTime,omitempty"`
	Endpoints          []*Endpoint `json:"endpoints,omitempty"`
}
