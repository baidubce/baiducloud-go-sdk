package bcc

type ChangeVpcRequest struct {
	InstanceId                 *string   `json:"instanceId,omitempty"`
	SubnetId                   *string   `json:"subnetId,omitempty"`
	InternalIp                 *string   `json:"internalIp,omitempty"`
	SecurityGroupIds           []*string `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []*string `json:"enterpriseSecurityGroupIds,omitempty"`
	Reboot                     *bool     `json:"reboot,omitempty"`
}
