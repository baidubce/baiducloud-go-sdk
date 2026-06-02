package pfs

type MountTargetListClientsRequest struct {
	Action        *string `json:"-"`
	MountTargetId *string `json:"mountTargetId,omitempty"`
	MaxKeys       *int32  `json:"maxKeys,omitempty"`
	Manner        *string `json:"manner,omitempty"`
	Marker        *string `json:"marker,omitempty"`
}
