package vpc

type RejectPeerConnRequest struct {
	PeerConnId  *string `json:"-"`
	ClientToken *string `json:"-"`
}
