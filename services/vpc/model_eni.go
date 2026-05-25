package vpc

type Eni struct {
	EniId            *string      `json:"eniId,omitempty"`
	Name             *string      `json:"name,omitempty"`
	ZoneName         *string      `json:"zoneName,omitempty"`
	Description      *string      `json:"description,omitempty"`
	InstanceId       *string      `json:"instanceId,omitempty"`
	MacAddress       *string      `json:"macAddress,omitempty"`
	VpcId            *string      `json:"vpcId,omitempty"`
	SubnetId         *string      `json:"subnetId,omitempty"`
	Status           *string      `json:"status,omitempty"`
	PrivateIpSet     []*PrivateIP `json:"privateIpSet,omitempty"`
	Ipv6PrivateIpSet []*PrivateIP `json:"ipv6PrivateIpSet,omitempty"`
	Tags             []*TagModel  `json:"tags,omitempty"`
}
