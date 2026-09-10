package cce

type CreateAShrinkingNodeGroupTaskV2Request struct {
	ClusterID            *string       `json:"-"`
	InstanceGroupID      *string       `json:"-"`
	InstancesToBeRemoved []*string     `json:"instancesToBeRemoved,omitempty"`
	K8sNodesToBeRemoved  []*string     `json:"k8sNodesToBeRemoved,omitempty"`
	CleanPolicy          *string       `json:"cleanPolicy,omitempty"`
	DeleteOption         *DeleteOption `json:"deleteOption,omitempty"`
}
