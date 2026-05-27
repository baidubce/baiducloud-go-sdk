package csn

type DeleteRegionBandwidthRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	LocalRegion *string `json:"localRegion,omitempty"`
	PeerRegion  *string `json:"peerRegion,omitempty"`
}
