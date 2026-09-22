package scs

type UpdateSecurityGroupRequest struct {
	InstanceId       *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
