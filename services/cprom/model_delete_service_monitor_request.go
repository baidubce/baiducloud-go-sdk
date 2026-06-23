package cprom

type DeleteServiceMonitorRequest struct {
	ServiceMonitorName *string `json:"-"`
	InstanceId         *string `json:"-"`
	AgentId            *string `json:"-"`
}
