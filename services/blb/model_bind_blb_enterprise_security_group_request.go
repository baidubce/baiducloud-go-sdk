package blb

type BindBlbEnterpriseSecurityGroupRequest struct {
	BlbId                      *string   `json:"-"`
	ClientToken                *string   `json:"-"`
	EnterpriseSecurityGroupIds []*string `json:"enterpriseSecurityGroupIds,omitempty"`
}
