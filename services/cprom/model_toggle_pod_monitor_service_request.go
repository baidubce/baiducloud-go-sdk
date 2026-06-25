package cprom

type TogglePodMonitorServiceRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"-"`
	AgentId    *string `json:"-"`
}
