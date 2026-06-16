package bcc

type RemoteCopyRequest struct {
	Name       *string `json:"name,omitempty"`
	DestRegion *string `json:"destRegion,omitempty"`
}
