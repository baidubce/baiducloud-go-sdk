package vpc

type UpdateEniEnterpriseSecurityGroupRequest struct {
	EniId                      *string   `json:"-"`
	ClientToken                *string   `json:"-"`
	EnterpriseSecurityGroupIds []*string `json:"enterpriseSecurityGroupIds,omitempty"`
}
