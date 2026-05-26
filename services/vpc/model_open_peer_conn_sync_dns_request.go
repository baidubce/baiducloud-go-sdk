package vpc

type OpenPeerConnSyncDnsRequest struct {
	PeerConnId  *string `json:"-"`
	Role        *string `json:"-"`
	ClientToken *string `json:"-"`
}
