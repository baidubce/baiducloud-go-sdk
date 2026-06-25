package cprom

type PodMonitorSpec struct {
	NamespaceSelector   *NamespaceSelector    `json:"namespaceSelector,omitempty"`
	PodMetricsEndpoints []*PodMetricsEndpoint `json:"podMetricsEndpoints,omitempty"`
	Selector            *LabelSelector        `json:"selector,omitempty"`
}
