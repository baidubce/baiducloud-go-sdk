package bcc

type EniInfo struct {
	EniId            *string      `json:"eniId,omitempty"`
	Name             *string      `json:"name,omitempty"`
	VpcId            *string      `json:"vpcId,omitempty"`
	SubnetId         *string      `json:"subnetId,omitempty"`
	ZoneName         *string      `json:"zoneName,omitempty"`
	Description      *string      `json:"description,omitempty"`
	CreatedTime      *string      `json:"createdTime,omitempty"`
	InstanceId       *string      `json:"instanceId,omitempty"`
	MacAddress       *string      `json:"macAddress,omitempty"`
	Status           *string      `json:"status,omitempty"`
	SecurityGroupIds []*string    `json:"securityGroupIds,omitempty"`
	PrivateIpSet     []*IpAddress `json:"privateIpSet,omitempty"`
}
