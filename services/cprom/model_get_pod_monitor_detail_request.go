package cprom

type GetPodMonitorDetailRequest struct {
	PodMonitorName *string `json:"-"`
	InstanceId     *string `json:"-"`
	AgentId        *string `json:"-"`
}
