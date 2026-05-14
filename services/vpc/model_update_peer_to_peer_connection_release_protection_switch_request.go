package vpc

type UpdatePeerToPeerConnectionReleaseProtectionSwitchRequest struct {
	PeerConnId    *string `json:"-"`
	DeleteProtect *bool   `json:"deleteProtect,omitempty"`
}
