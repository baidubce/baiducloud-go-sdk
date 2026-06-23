package cprom

type DeletePodmonitorRequest struct {
	PodMonitorName *string `json:"-"`
	InstanceId     *string `json:"-"`
	AgentId        *string `json:"-"`
}
