package vpc

type EnablePeerToPeerConnectionToSynchronizeDnsRequest struct {
	PeerConnId  *string `json:"-"`
	Role        *string `json:"-"`
	ClientToken *string `json:"-"`
}
