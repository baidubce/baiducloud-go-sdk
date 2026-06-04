package privatezone

type VpcRegion struct {
	Region *string   `json:"region,omitempty"`
	VpcIds []*string `json:"vpcIds,omitempty"`
}
