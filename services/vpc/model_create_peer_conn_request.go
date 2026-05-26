package vpc

type CreatePeerConnRequest struct {
	ClientToken     *string     `json:"-"`
	BandwidthInMbps *int32      `json:"bandwidthInMbps,omitempty"`
	Description     *string     `json:"description,omitempty"`
	LocalIfName     *string     `json:"localIfName,omitempty"`
	LocalVpcId      *string     `json:"localVpcId,omitempty"`
	PeerAccountId   *string     `json:"peerAccountId,omitempty"`
	PeerVpcId       *string     `json:"peerVpcId,omitempty"`
	PeerRegion      *string     `json:"peerRegion,omitempty"`
	PeerIfName      *string     `json:"peerIfName,omitempty"`
	Billing         *Billing    `json:"billing,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	ResourceGroupId *string     `json:"resourceGroupId,omitempty"`
	DeleteProtect   *bool       `json:"deleteProtect,omitempty"`
}
