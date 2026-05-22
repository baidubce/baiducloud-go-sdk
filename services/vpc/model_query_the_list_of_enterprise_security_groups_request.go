package vpc

type QueryTheListOfEnterpriseSecurityGroupsRequest struct {
	Marker     *string `json:"-"`
	MaxKeys    *int32  `json:"-"`
	InstanceId *string `json:"-"`
}
