package cprom

type UpdatePodMonitorRequest struct {
	PodMonitorName *string     `json:"-"`
	InstanceId     *string     `json:"-"`
	AgentId        *string     `json:"-"`
	Enable         *string     `json:"enable,omitempty"`
	PodMonitor     *PodMonitor `json:"podMonitor,omitempty"`
}
