package vpc

type UpdatePeerConnDeleteProtectRequest struct {
	PeerConnId    *string `json:"-"`
	ClientToken   *string `json:"-"`
	DeleteProtect *bool   `json:"deleteProtect,omitempty"`
}
