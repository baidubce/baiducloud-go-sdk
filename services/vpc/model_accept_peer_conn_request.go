package vpc

type AcceptPeerConnRequest struct {
	PeerConnId  *string `json:"-"`
	ClientToken *string `json:"-"`
}
