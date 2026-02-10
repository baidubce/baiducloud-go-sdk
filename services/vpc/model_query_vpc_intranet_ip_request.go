package vpc

type QueryVpcIntranetIpRequest struct {
	VpcId              *string   `json:"-"`
	PrivateIpAddresses []*string `json:"-"`
	PrivateIpRange     *string   `json:"-"`
}
