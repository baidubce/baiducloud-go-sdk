package rds

type SimpleVpcVo struct {
	VpcId         *string   `json:"vpcId,omitempty"`
	ShortId       *string   `json:"shortId,omitempty"`
	Name          *string   `json:"name,omitempty"`
	Cidr          *string   `json:"cidr,omitempty"`
	Status        *int32    `json:"status,omitempty"`
	CreateTime    *string   `json:"createTime,omitempty"`
	Description   *string   `json:"description,omitempty"`
	DefaultVpc    *bool     `json:"defaultVpc,omitempty"`
	Ipv6Cidr      *string   `json:"ipv6Cidr,omitempty"`
	AuxiliaryCidr []*string `json:"auxiliaryCidr,omitempty"`
	Relay         *bool     `json:"relay,omitempty"`
}
