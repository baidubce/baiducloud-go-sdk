package pfs

type ListL2PolicyRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	MaxKeys    *int32  `json:"maxKeys,omitempty"`
	Marker     *string `json:"marker,omitempty"`
}
