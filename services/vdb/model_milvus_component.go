package vdb

type MilvusComponent struct {
	DiskSizeType *string `json:"diskSizeType,omitempty"`
	Replicas     *int32  `json:"replicas,omitempty"`
	Spec         *string `json:"spec,omitempty"`
	VdbType      *string `json:"type,omitempty"`
}
