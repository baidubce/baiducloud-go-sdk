package scs

type UnbindSecurityGroupRequest struct {
	InstanceId       *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
