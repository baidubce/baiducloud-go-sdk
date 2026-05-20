package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheDetailsOfTheDedicatedGatewayResponse struct {
	bce.BaseResponse
	EtGatewayId         *string     `json:"etGatewayId,omitempty"`
	Name                *string     `json:"name,omitempty"`
	Status              *string     `json:"status,omitempty"`
	Speed               *int32      `json:"speed,omitempty"`
	CreateTime          *string     `json:"createTime,omitempty"`
	Description         *string     `json:"description,omitempty"`
	VpcId               *string     `json:"vpcId,omitempty"`
	EtId                *string     `json:"etId,omitempty"`
	ChannelId           *string     `json:"channelId,omitempty"`
	LocalCidrs          []*string   `json:"localCidrs,omitempty"`
	EnableIpv6          *int32      `json:"enableIpv6,omitempty"`
	Ipv6LocalCidrs      []*string   `json:"ipv6LocalCidrs,omitempty"`
	HealthCheckSourceIp *string     `json:"healthCheckSourceIp,omitempty"`
	HealthCheckDestIp   *string     `json:"healthCheckDestIp,omitempty"`
	HealthCheckInterval *int32      `json:"healthCheckInterval,omitempty"`
	HealthThreshold     *int32      `json:"healthThreshold,omitempty"`
	UnhealthThreshold   *int32      `json:"unhealthThreshold,omitempty"`
	HealthCheckType     *string     `json:"healthCheckType,omitempty"`
	HealthCheckPort     *int32      `json:"healthCheckPort,omitempty"`
	Tags                []*TagModel `json:"tags,omitempty"`
}
