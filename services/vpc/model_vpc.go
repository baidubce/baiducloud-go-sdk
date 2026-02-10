package vpc

type Vpc struct {
	VpcId         *string     `json:"vpcId,omitempty"`
	Name          *string     `json:"name,omitempty"`
	Cidr          *string     `json:"cidr,omitempty"`
	Ipv6Cidr      *string     `json:"ipv6Cidr,omitempty"`
	Description   *string     `json:"description,omitempty"`
	IsDefault     *bool       `json:"isDefault,omitempty"`
	Relay         *bool       `json:"relay,omitempty"`
	SecondaryCidr []*string   `json:"secondaryCidr,omitempty"`
	Tags          []*TagModel `json:"tags,omitempty"`
}
