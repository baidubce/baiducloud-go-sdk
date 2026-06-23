package cprom

type ToggleServiceMonitorServiceRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"-"`
	AgentId    *string `json:"-"`
}
