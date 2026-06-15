package bcc

type BatchChangeToPrepaidRequest struct {
	Config []*PrepayConfig `json:"config,omitempty"`
}
