package vpc

type DeleteSecurityGroupRequest struct {
	SecurityGroupId *string `json:"-"`
	ClientToken     *string `json:"-"`
}
