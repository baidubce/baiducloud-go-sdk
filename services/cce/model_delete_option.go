package cce

type DeleteOption struct {
	DeleteResource    *bool `json:"deleteResource,omitempty"`
	DeleteCDSSnapshot *bool `json:"deleteCDSSnapshot,omitempty"`
	MoveOut           *bool `json:"moveOut,omitempty"`
}
