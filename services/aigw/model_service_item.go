package aigw

type ServiceItem struct {
	ServiceName *string   `json:"serviceName,omitempty"`
	ClusterIds  []*string `json:"clusterIds,omitempty"`
}
