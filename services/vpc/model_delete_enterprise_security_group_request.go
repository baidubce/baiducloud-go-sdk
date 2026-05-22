package vpc

type DeleteEnterpriseSecurityGroupRequest struct {
	EnterpriseSecurityGroupId *string `json:"-"`
	ClientToken               *string `json:"-"`
}
