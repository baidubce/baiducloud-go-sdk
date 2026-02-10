package vpc

type CloseVpcRelayRequest struct {
	VpcId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
