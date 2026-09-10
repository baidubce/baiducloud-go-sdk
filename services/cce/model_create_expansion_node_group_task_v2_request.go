package cce

type CreateExpansionNodeGroupTaskV2Request struct {
	ClusterID       *string `json:"-"`
	InstanceGroupID *string `json:"-"`
	UpToReplicas    *int32  `json:"-"`
	UpReplicas      *int32  `json:"-"`
}
