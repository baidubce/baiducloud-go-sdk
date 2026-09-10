package cce

type ModifyIGAutoScalerRequest struct {
	ClusterID            *string `json:"-"`
	InstanceGroupID      *string `json:"-"`
	Enabled              *bool   `json:"enabled,omitempty"`
	MinReplicas          *int32  `json:"minReplicas,omitempty"`
	MaxReplicas          *int32  `json:"maxReplicas,omitempty"`
	ScalingGroupPriority *int32  `json:"scalingGroupPriority,omitempty"`
}
