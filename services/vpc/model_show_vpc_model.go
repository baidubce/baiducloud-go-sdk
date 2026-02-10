package vpc

type ShowVpcModel struct {
	VpcId         *string     `json:"vpcId,omitempty"`
	Name          *string     `json:"name,omitempty"`
	Cidr          *string     `json:"cidr,omitempty"`
	Ipv6Cidr      *string     `json:"ipv6Cidr,omitempty"`
	Description   *string     `json:"description,omitempty"`
	IsDefault     *bool       `json:"isDefault,omitempty"`
	Relay         *bool       `json:"relay,omitempty"`
	Subnets       []*Subnet   `json:"subnets,omitempty"`
	SecondaryCidr []*string   `json:"secondaryCidr,omitempty"`
	Tags          []*TagModel `json:"tags,omitempty"`
}
