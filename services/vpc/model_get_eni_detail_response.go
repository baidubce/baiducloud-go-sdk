package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetEniDetailResponse struct {
	bce.BaseResponse
	EniId                      *string      `json:"eniId,omitempty"`
	Name                       *string      `json:"name,omitempty"`
	Description                *string      `json:"description,omitempty"`
	VpcId                      *string      `json:"vpcId,omitempty"`
	SubnetId                   *string      `json:"subnetId,omitempty"`
	MacAddress                 *string      `json:"macAddress,omitempty"`
	Status                     *string      `json:"status,omitempty"`
	ZoneName                   *string      `json:"zoneName,omitempty"`
	InstanceId                 *string      `json:"instanceId,omitempty"`
	PrivateIpSet               []*PrivateIP `json:"privateIpSet,omitempty"`
	Ipv6PrivateIpSet           []*PrivateIP `json:"ipv6PrivateIpSet,omitempty"`
	SecurityGroupIds           []*string    `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []*string    `json:"enterpriseSecurityGroupIds,omitempty"`
	CreatedTime                *string      `json:"createdTime,omitempty"`
	Tags                       []*TagModel  `json:"tags,omitempty"`
}
