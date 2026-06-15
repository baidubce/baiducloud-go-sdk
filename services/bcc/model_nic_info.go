package bcc

type NicInfo struct {
	EniId                    *string    `json:"eniId,omitempty"`
	EniUuid                  *string    `json:"eniUuid,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	BccType                  *string    `json:"type,omitempty"`
	SubnetId                 *string    `json:"subnetId,omitempty"`
	SubnetType               *string    `json:"subnetType,omitempty"`
	Az                       *string    `json:"az,omitempty"`
	Description              *string    `json:"description,omitempty"`
	DeviceId                 *string    `json:"deviceId,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	MacAddress               *string    `json:"macAddress,omitempty"`
	VpcId                    *string    `json:"vpcId,omitempty"`
	CreatedTime              *string    `json:"createdTime,omitempty"`
	EniNum                   *int32     `json:"eniNum,omitempty"`
	EriNum                   *int32     `json:"eriNum,omitempty"`
	EriInfos                 []*EriInfo `json:"eriInfos,omitempty"`
	EniInfos                 []*EniInfo `json:"eniInfos,omitempty"`
	Ips                      []*IpInfo  `json:"ips,omitempty"`
	Ipv6s                    []*IpInfo  `json:"ipv6s,omitempty"`
	SecurityGroups           []*string  `json:"securityGroups,omitempty"`
	EnterpriseSecurityGroups []*string  `json:"enterpriseSecurityGroups,omitempty"`
}
