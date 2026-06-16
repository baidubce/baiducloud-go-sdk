package bcc

type ListSecurityGroupsRequest struct {
	Marker           *string `json:"-"`
	MaxKeys          *int32  `json:"-"`
	InstanceId       *string `json:"-"`
	VpcId            *string `json:"-"`
	SecurityGroupId  *string `json:"-"`
	SecurityGroupIds *string `json:"-"`
}
