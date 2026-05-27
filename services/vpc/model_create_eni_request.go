package vpc

type CreateEniRequest struct {
	ClientToken                 *string      `json:"-"`
	Name                        *string      `json:"name,omitempty"`
	SubnetId                    *string      `json:"subnetId,omitempty"`
	SecurityGroupIds            []*string    `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds  []*string    `json:"enterpriseSecurityGroupIds,omitempty"`
	PrivateIpSet                []*PrivateIP `json:"privateIpSet,omitempty"`
	Ipv6PrivateIpSet            []*PrivateIP `json:"ipv6PrivateIpSet,omitempty"`
	Description                 *string      `json:"description,omitempty"`
	NetworkInterfaceTrafficMode *string      `json:"networkInterfaceTrafficMode,omitempty"`
}
