package blb

type BindBlbSecurityGroupRequest struct {
	BlbId            *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
