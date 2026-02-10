package vpc

type QueryTheIpAddressesOccupiedByProductsWithinVpcRequest struct {
	VpcId        *string `json:"-"`
	SubnetId     *string `json:"-"`
	ResourceType *string `json:"-"`
	PageNo       *int32  `json:"-"`
	PageSize     *int32  `json:"-"`
}
