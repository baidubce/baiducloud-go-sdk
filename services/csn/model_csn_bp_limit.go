package csn

type CsnBpLimit struct {
	CsnBpId     *string `json:"csnBpId,omitempty"`
	CsnId       *string `json:"csnId,omitempty"`
	LocalRegion *string `json:"localRegion,omitempty"`
	PeerRegion  *string `json:"peerRegion,omitempty"`
	Bandwidth   *int32  `json:"bandwidth,omitempty"`
}
