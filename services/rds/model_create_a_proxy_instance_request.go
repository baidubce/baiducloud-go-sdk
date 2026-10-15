package rds

type CreateAProxyInstanceRequest struct {
	ClientToken      *string      `json:"-"`
	Billing          *Billing     `json:"billing,omitempty"`
	SourceInstanceId *string      `json:"sourceInstanceId,omitempty"`
	InstanceName     *string      `json:"instanceName,omitempty"`
	NodeAmount       *int32       `json:"nodeAmount,omitempty"`
	ZoneNames        []*string    `json:"zoneNames,omitempty"`
	VpcId            *string      `json:"vpcId,omitempty"`
	IsDirectPay      *bool        `json:"isDirectPay,omitempty"`
	Subnets          []*SubnetMap `json:"subnets,omitempty"`
	Ovip             *string      `json:"ovip,omitempty"`
	EntryPort        *int32       `json:"entryPort,omitempty"`
	ResourceGroupId  *string      `json:"resourceGroupId,omitempty"`
}
