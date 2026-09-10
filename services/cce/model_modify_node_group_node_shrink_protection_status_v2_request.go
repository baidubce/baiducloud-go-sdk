package cce

type ModifyNodeGroupNodeShrinkProtectionStatusV2Request struct {
	ClusterID         *string   `json:"-"`
	InstanceIDs       []*string `json:"instanceIDs,omitempty"`
	ScaleDownDisabled *bool     `json:"scaleDownDisabled,omitempty"`
}
