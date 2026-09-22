package scs

type VpcInfo struct {
	VpcId   *string `json:"vpcId,omitempty"`
	VpcName *string `json:"vpcName,omitempty"`
	VpcCidr *string `json:"vpcCidr,omitempty"`
}
