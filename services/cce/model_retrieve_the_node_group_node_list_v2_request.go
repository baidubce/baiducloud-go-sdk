package cce

type RetrieveTheNodeGroupNodeListV2Request struct {
	ClusterID       *string `json:"-"`
	InstanceGroupID *string `json:"-"`
	PageNo          *int32  `json:"-"`
	PageSize        *int32  `json:"-"`
}
