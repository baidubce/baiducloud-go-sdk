package privatezone

type DisassociateVpcRequest struct {
	ZoneId      *string   `json:"-"`
	Action      *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Region      *string   `json:"region,omitempty"`
	VpcIds      []*string `json:"vpcIds,omitempty"`
}
