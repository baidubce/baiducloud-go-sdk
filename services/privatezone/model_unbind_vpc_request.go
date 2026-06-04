package privatezone

type UnbindVpcRequest struct {
	ZoneId      *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Region      *string   `json:"region,omitempty"`
	VpcIds      []*string `json:"vpcIds,omitempty"`
}
