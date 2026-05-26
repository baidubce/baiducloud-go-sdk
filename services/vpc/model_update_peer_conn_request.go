package vpc

type UpdatePeerConnRequest struct {
	PeerConnId  *string `json:"-"`
	ClientToken *string `json:"-"`
	LocalIfId   *string `json:"localIfId,omitempty"`
	Description *string `json:"description,omitempty"`
	LocalIfName *string `json:"localIfName,omitempty"`
}
