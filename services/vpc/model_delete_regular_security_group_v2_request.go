package vpc

type DeleteRegularSecurityGroupV2Request struct {
	SecurityGroupId *string `json:"-"`
	ClientToken     *string `json:"-"`
}
