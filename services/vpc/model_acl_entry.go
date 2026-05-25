package vpc

type AclEntry struct {
	SubnetId   *string    `json:"subnetId,omitempty"`
	SubnetName *string    `json:"subnetName,omitempty"`
	SubnetCidr *string    `json:"subnetCidr,omitempty"`
	AclRules   []*AclRule `json:"aclRules,omitempty"`
}
