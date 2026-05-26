package vpc

type ClosePeerConnSyncDnsRequest struct {
	PeerConnId  *string `json:"-"`
	Role        *string `json:"-"`
	ClientToken *string `json:"-"`
}
