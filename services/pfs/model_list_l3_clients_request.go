package pfs

type ListL3ClientsRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	Marker     *string `json:"marker,omitempty"`
	MaxKeys    *int32  `json:"maxKeys,omitempty"`
}
