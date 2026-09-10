package cce

type CAConfig struct {
	ReplicaCount                     *int32             `json:"replicaCount,omitempty"`
	ScaleDownEnabled                 *bool              `json:"scaleDownEnabled,omitempty"`
	ScaleDownUtilizationThreshold    *int32             `json:"scaleDownUtilizationThreshold,omitempty"`
	ScaleDownGPUUtilizationThreshold *int32             `json:"scaleDownGPUUtilizationThreshold,omitempty"`
	ScaleDownUnneededTime            *int32             `json:"scaleDownUnneededTime,omitempty"`
	ScaleDownDelayAfterAdd           *int32             `json:"scaleDownDelayAfterAdd,omitempty"`
	MaxEmptyBulkDelete               *int32             `json:"maxEmptyBulkDelete,omitempty"`
	SkipNodesWithLocalStorage        *bool              `json:"skipNodesWithLocalStorage,omitempty"`
	SkipNodesWithSystemPods          *bool              `json:"skipNodesWithSystemPods,omitempty"`
	Expander                         *string            `json:"expander,omitempty"`
	CustomConfigs                    *map[string]string `json:"customConfigs,omitempty"`
}
