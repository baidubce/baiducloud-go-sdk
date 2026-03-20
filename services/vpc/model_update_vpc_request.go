package vpc

type UpdateVpcRequest struct {
	VpcId         *string   `json:"-"`
	ClientToken   *string   `json:"-"`
	Name          *string   `json:"name,omitempty"`
	Description   *string   `json:"description,omitempty"`
	EnableIpv6    *bool     `json:"enableIpv6,omitempty"`
	SecondaryCidr []*string `json:"secondaryCidr,omitempty"`
}
