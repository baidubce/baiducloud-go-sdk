package pfs

type ListL2BucketLinkRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	MaxKeys    *int32  `json:"maxKeys,omitempty"`
	Manner     *string `json:"manner,omitempty"`
	Marker     *string `json:"marker,omitempty"`
}
