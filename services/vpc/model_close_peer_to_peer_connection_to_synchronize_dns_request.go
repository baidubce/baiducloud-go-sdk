package vpc

type ClosePeerToPeerConnectionToSynchronizeDnsRequest struct {
	PeerConnId  *string `json:"-"`
	Role        *string `json:"-"`
	ClientToken *string `json:"-"`
}
