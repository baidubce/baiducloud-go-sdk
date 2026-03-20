package vpc

type CreateVpcRequest struct {
	ClientToken *string     `json:"-"`
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Cidr        *string     `json:"cidr,omitempty"`
	EnableIpv6  *bool       `json:"enableIpv6,omitempty"`
	Tags        []*TagModel `json:"tags,omitempty"`
}
