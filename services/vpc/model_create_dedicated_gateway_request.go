package vpc

type CreateDedicatedGatewayRequest struct {
	ClientToken     *string     `json:"-"`
	Name            *string     `json:"name,omitempty"`
	VpcId           *string     `json:"vpcId,omitempty"`
	Speed           *int32      `json:"speed,omitempty"`
	Description     *string     `json:"description,omitempty"`
	EtId            *string     `json:"etId,omitempty"`
	ChannelId       *string     `json:"channelId,omitempty"`
	LocalCidrs      []*string   `json:"localCidrs,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	ResourceGroupId *string     `json:"resourceGroupId,omitempty"`
}
