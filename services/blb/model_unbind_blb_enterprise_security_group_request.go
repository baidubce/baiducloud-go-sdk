package blb

type UnbindBlbEnterpriseSecurityGroupRequest struct {
	BlbId                      *string   `json:"-"`
	ClientToken                *string   `json:"-"`
	EnterpriseSecurityGroupIds []*string `json:"enterpriseSecurityGroupIds,omitempty"`
}
