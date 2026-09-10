package cce

type ClusterAutoscalerSpec struct {
	Enabled              *bool  `json:"enabled,omitempty"`
	MinReplicas          *int32 `json:"minReplicas,omitempty"`
	MaxReplicas          *int32 `json:"maxReplicas,omitempty"`
	ScalingGroupPriority *int32 `json:"scalingGroupPriority,omitempty"`
}
