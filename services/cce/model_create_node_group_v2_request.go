package cce

type CreateNodeGroupV2Request struct {
	ClusterID             *string                `json:"-"`
	InstanceGroupName     *string                `json:"instanceGroupName,omitempty"`
	ClusterRole           *string                `json:"clusterRole,omitempty"`
	ShrinkPolicy          *string                `json:"shrinkPolicy,omitempty"`
	UpdatePolicy          *string                `json:"updatePolicy,omitempty"`
	CleanPolicy           *string                `json:"cleanPolicy,omitempty"`
	InstanceTemplate      *InstanceTemplate      `json:"instanceTemplate,omitempty"`
	Replicas              *int32                 `json:"replicas,omitempty"`
	ClusterAutoscalerSpec *ClusterAutoscalerSpec `json:"clusterAutoscalerSpec,omitempty"`
}
