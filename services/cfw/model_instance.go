package cfw

type Instance struct {
	InstanceId      *string `json:"instanceId,omitempty"`
	InstanceName    *string `json:"instanceName,omitempty"`
	Status          *string `json:"status,omitempty"`
	Region          *string `json:"region,omitempty"`
	CfwId           *string `json:"cfwId,omitempty"`
	CfwName         *string `json:"cfwName,omitempty"`
	VpcId           *string `json:"vpcId,omitempty"`
	VpcName         *string `json:"vpcName,omitempty"`
	PublicIp        *string `json:"publicIp,omitempty"`
	Role            *string `json:"role,omitempty"`
	LocalIfId       *string `json:"localIfId,omitempty"`
	LocalIfName     *string `json:"localIfName,omitempty"`
	PeerRegion      *string `json:"peerRegion,omitempty"`
	PeerVpcId       *string `json:"peerVpcId,omitempty"`
	PeerVpcName     *string `json:"peerVpcName,omitempty"`
	MemberId        *string `json:"memberId,omitempty"`
	MemberName      *string `json:"memberName,omitempty"`
	MemberAccountId *string `json:"memberAccountId,omitempty"`
}
