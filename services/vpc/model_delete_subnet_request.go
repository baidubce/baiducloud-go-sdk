package vpc

type DeleteSubnetRequest struct {
	SubnetId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
