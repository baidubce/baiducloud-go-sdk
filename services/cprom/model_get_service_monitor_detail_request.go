package cprom

type GetServiceMonitorDetailRequest struct {
	ServiceMonitorName *string `json:"-"`
	InstanceId         *string `json:"-"`
	AgentId            *string `json:"-"`
}
