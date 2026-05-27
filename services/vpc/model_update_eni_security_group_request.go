package vpc

type UpdateEniSecurityGroupRequest struct {
	EniId            *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty"`
}
