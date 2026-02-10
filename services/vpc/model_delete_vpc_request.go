package vpc

type DeleteVpcRequest struct {
	VpcId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
