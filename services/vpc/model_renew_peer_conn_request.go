package vpc

type RenewPeerConnRequest struct {
	PeerConnId  *string  `json:"-"`
	ClientToken *string  `json:"-"`
	Billing     *Billing `json:"billing,omitempty"`
}
