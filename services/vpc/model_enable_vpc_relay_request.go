package vpc

type EnableVpcRelayRequest struct {
	VpcId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
