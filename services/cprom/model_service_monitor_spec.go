package cprom

type ServiceMonitorSpec struct {
	Endpoints         []*ServiceMonitorEndpoint `json:"endpoints,omitempty"`
	NamespaceSelector *NamespaceSelector        `json:"namespaceSelector,omitempty"`
	Selector          *LabelSelector            `json:"selector,omitempty"`
}
