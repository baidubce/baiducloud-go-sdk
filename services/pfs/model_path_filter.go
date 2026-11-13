package pfs

type PathFilter struct {
	PfsType *int32  `json:"type,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
}
