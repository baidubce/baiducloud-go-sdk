package scs

type BindSecurityGroupRequest struct {
	InstanceId       *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
