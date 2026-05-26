package vpc

type PeerConn struct {
	PeerConnId      *string `json:"peerConnId,omitempty"`
	Role            *string `json:"role,omitempty"`
	Status          *string `json:"status,omitempty"`
	BandwidthInMbps *string `json:"bandwidthInMbps,omitempty"`
	Description     *string `json:"description,omitempty"`
	LocalIfId       *string `json:"localIfId,omitempty"`
	LocalIfName     *string `json:"localIfName,omitempty"`
	LocalVpcId      *string `json:"localVpcId,omitempty"`
	LocalRegion     *string `json:"localRegion,omitempty"`
	PeerVpcId       *string `json:"peerVpcId,omitempty"`
	PeerRegion      *string `json:"peerRegion,omitempty"`
	PeerAccountId   *string `json:"peerAccountId,omitempty"`
	PaymentTiming   *string `json:"paymentTiming,omitempty"`
	DnsStatus       *string `json:"dnsStatus,omitempty"`
	CreatedTime     *string `json:"createdTime,omitempty"`
	ExpiredTime     *string `json:"expiredTime,omitempty"`
}
