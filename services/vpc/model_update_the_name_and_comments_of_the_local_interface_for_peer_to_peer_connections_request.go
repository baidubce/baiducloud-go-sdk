package vpc

type UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsRequest struct {
	PeerConnId  *string `json:"-"`
	LocalIfId   *string `json:"localIfId,omitempty"`
	Description *string `json:"description,omitempty"`
	LocalIfName *string `json:"localIfName,omitempty"`
}
