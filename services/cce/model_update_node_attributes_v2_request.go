package cce

type UpdateNodeAttributesV2Request struct {
	ClusterID           *string            `json:"-"`
	InstanceID          *string            `json:"-"`
	Labels              *map[string]string `json:"labels,omitempty"`
	Annotations         *map[string]string `json:"annotations,omitempty"`
	Taints              []*Taint           `json:"taints,omitempty"`
	CceInstancePriority *int32             `json:"cceInstancePriority,omitempty"`
}
