package bcc

type BindInstanceToSecurityGroupRequest struct {
	InstanceId      *string `json:"-"`
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
}
