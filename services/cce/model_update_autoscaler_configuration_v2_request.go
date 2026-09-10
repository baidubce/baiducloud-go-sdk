package cce

type UpdateAutoscalerConfigurationV2Request struct {
	ClusterID                        *string                   `json:"-"`
	Expander                         *string                   `json:"expander,omitempty"`
	InstanceGroups                   []*map[string]interface{} `json:"instanceGroups,omitempty"`
	KubeVersion                      *string                   `json:"kubeVersion,omitempty"`
	MaxEmptyBulkDelete               *int32                    `json:"maxEmptyBulkDelete,omitempty"`
	ScaleDownDelayAfterAdd           *int32                    `json:"scaleDownDelayAfterAdd,omitempty"`
	ScaleDownEnabled                 *bool                     `json:"scaleDownEnabled,omitempty"`
	ScaleDownGPUUtilizationThreshold *int32                    `json:"scaleDownGPUUtilizationThreshold,omitempty"`
	ScaleDownUnneededTime            *int32                    `json:"scaleDownUnneededTime,omitempty"`
	ScaleDownUtilizationThreshold    *int32                    `json:"scaleDownUtilizationThreshold,omitempty"`
	SkipNodesWithLocalStorage        *bool                     `json:"skipNodesWithLocalStorage,omitempty"`
	SkipNodesWithSystemPods          *bool                     `json:"skipNodesWithSystemPods,omitempty"`
	CustomConfigs                    *map[string]string        `json:"customConfigs,omitempty"`
}
