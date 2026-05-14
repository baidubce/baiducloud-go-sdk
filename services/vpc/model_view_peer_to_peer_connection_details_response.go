package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ViewPeerToPeerConnectionDetailsResponse struct {
	bce.BaseResponse
	PeerConnId      *string         `json:"peerConnId,omitempty"`
	Role            *string         `json:"role,omitempty"`
	Status          *PeerConnStatus `json:"status,omitempty"`
	BandwidthInMbps *int32          `json:"bandwidthInMbps,omitempty"`
	Description     *string         `json:"description,omitempty"`
	LocalIfId       *string         `json:"localIfId,omitempty"`
	LocalIfName     *string         `json:"localIfName,omitempty"`
	LocalVpcId      *string         `json:"localVpcId,omitempty"`
	LocalRegion     *string         `json:"localRegion,omitempty"`
	PeerVpcId       *string         `json:"peerVpcId,omitempty"`
	PeerRegion      *string         `json:"peerRegion,omitempty"`
	PeerAccountId   *string         `json:"peerAccountId,omitempty"`
	CreatedTime     *string         `json:"createdTime,omitempty"`
	ExpiredTime     *string         `json:"expiredTime,omitempty"`
	DnsStatus       *DnsStatus      `json:"dnsStatus,omitempty"`
	PaymentTiming   *string         `json:"paymentTiming,omitempty"`
	Tags            []*TagModel     `json:"tags,omitempty"`
	DeleteProtect   *bool           `json:"deleteProtect,omitempty"`
}
