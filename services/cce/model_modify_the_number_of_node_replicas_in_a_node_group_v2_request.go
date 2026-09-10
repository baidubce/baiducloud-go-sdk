package cce

type ModifyTheNumberOfNodeReplicasInANodeGroupV2Request struct {
	ClusterID       *string       `json:"-"`
	InstanceGroupID *string       `json:"-"`
	Replicas        *int32        `json:"replicas,omitempty"`
	InstanceIDs     []*string     `json:"instanceIDs,omitempty"`
	DeleteInstance  *bool         `json:"deleteInstance,omitempty"`
	DeleteOption    *DeleteOption `json:"deleteOption,omitempty"`
}
