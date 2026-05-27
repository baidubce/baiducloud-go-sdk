package vpc

type QueryEnterpriseSecurityGroupListRequest struct {
	Marker     *string `json:"-"`
	MaxKeys    *int32  `json:"-"`
	InstanceId *string `json:"-"`
}
