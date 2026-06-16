package bcc

type BindInstanceSecurityGroupRequest struct {
	InstanceIds       []*string `json:"instanceIds,omitempty"`
	SecurityGroupIds  []*string `json:"securityGroupIds,omitempty"`
	SecurityGroupType *string   `json:"securityGroupType,omitempty"`
}
