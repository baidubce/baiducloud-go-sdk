package vpc

type EtGateway struct {
	EtGatewayId    *string     `json:"etGatewayId,omitempty"`
	Name           *string     `json:"name,omitempty"`
	Status         *string     `json:"status,omitempty"`
	Speed          *int32      `json:"speed,omitempty"`
	CreateTime     *string     `json:"createTime,omitempty"`
	Description    *string     `json:"description,omitempty"`
	VpcId          *string     `json:"vpcId,omitempty"`
	EtId           *string     `json:"etId,omitempty"`
	ChannelId      *string     `json:"channelId,omitempty"`
	LocalCidrs     []*string   `json:"localCidrs,omitempty"`
	EnableIpv6     *int32      `json:"enableIpv6,omitempty"`
	Ipv6LocalCidrs []*string   `json:"ipv6LocalCidrs,omitempty"`
	Tags           []*TagModel `json:"tags,omitempty"`
}
