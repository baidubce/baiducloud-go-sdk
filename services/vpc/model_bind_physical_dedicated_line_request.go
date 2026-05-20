package vpc

type BindPhysicalDedicatedLineRequest struct {
	EtGatewayId *string   `json:"-"`
	ClientToken *string   `json:"-"`
	EtId        *string   `json:"etId,omitempty"`
	ChannelId   *string   `json:"channelId,omitempty"`
	LocalCidrs  []*string `json:"localCidrs,omitempty"`
}
