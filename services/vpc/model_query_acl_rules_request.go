package vpc

type QueryAclRulesRequest struct {
	Marker   *string `json:"-"`
	MaxKeys  *int32  `json:"-"`
	SubnetId *string `json:"-"`
}
