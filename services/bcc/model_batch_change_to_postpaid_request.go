package bcc

type BatchChangeToPostpaidRequest struct {
	Config []*PostpayConfig `json:"config,omitempty"`
}
