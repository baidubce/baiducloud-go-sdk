package bcc

type UnbindInstanceFromSecurityGroupRequest struct {
	InstanceId      *string `json:"-"`
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
}
