package rapidfs

type DescribeCacheDeployGroupsRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	MaxKeys    *int32  `json:"maxKeys,omitempty"`
	Marker     *string `json:"marker,omitempty"`
}
