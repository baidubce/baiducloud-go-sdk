package aigw

type AihcArgs struct {
	AccountId        *string `json:"accountId,omitempty"`
	SubnetId         *string `json:"subnetId,omitempty"`
	SecurityGroupIds *string `json:"securityGroupIds,omitempty"`
	VpcCidr          *string `json:"vpcCidr,omitempty"`
	DomainPrefix     *string `json:"domainPrefix,omitempty"`
}
