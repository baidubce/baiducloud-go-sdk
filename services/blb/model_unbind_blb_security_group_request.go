package blb

type UnbindBlbSecurityGroupRequest struct {
	BlbId            *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
