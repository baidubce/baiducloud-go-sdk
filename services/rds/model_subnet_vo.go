package rds

type SubnetVo struct {
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
	SubnetId      *string   `json:"subnetId,omitempty"`
	Az            *string   `json:"az,omitempty"`
	VpcShortId    *string   `json:"vpcShortId,omitempty"`
	SubnetUuid    *string   `json:"subnetUuid,omitempty"`
	AccountId     *string   `json:"accountId,omitempty"`
	SubnetType    *int32    `json:"subnetType,omitempty"`
	RdsType       *int32    `json:"type,omitempty"`
	CreatedTime   *string   `json:"createdTime,omitempty"`
	UpdatedTime   *string   `json:"updatedTime,omitempty"`
	UsedIps       *int32    `json:"usedIps,omitempty"`
	TotalIps      *int32    `json:"totalIps,omitempty"`
	Tags          []*Tag    `json:"tags,omitempty"`
	ZoneName      *string   `json:"zoneName,omitempty"`
	EnableIpv6    *string   `json:"enable_ipv6,omitempty"`
}
