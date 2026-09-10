package cce

type DeleteNodesClusterScalingV2Request struct {
	ClusterID    *string       `json:"-"`
	DeleteOption *DeleteOption `json:"deleteOption,omitempty"`
	InstanceIDs  []*string     `json:"instanceIDs,omitempty"`
	ScaleDown    *bool         `json:"scaleDown,omitempty"`
}
