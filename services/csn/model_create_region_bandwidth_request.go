package csn

type CreateRegionBandwidthRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	LocalRegion *string `json:"localRegion,omitempty"`
	PeerRegion  *string `json:"peerRegion,omitempty"`
	Bandwidth   *int32  `json:"bandwidth,omitempty"`
}
