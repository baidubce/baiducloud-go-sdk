package vpc

type PeerToPeerConnectionRenewalRequest struct {
	PeerConnId  *string  `json:"-"`
	ClientToken *string  `json:"-"`
	Billing     *Billing `json:"billing,omitempty"`
}
