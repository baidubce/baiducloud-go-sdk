package cce

type DeleteNodeGroupV2Request struct {
	ClusterID          *string `json:"-"`
	InstanceGroupID    *string `json:"-"`
	DeleteInstances    *bool   `json:"-"`
	ReleaseAllResource *bool   `json:"-"`
}
