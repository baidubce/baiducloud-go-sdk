package cprom

type UpdateServiceMonitorRequest struct {
	ServiceMonitorName *string         `json:"-"`
	InstanceId         *string         `json:"-"`
	AgentId            *string         `json:"-"`
	Enable             *string         `json:"enable,omitempty"`
	ServiceMonitor     *ServiceMonitor `json:"serviceMonitor,omitempty"`
}
