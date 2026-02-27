package vpc

type OpenVpcRelayRequest struct {
	VpcId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
