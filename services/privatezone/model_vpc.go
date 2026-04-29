package privatezone

type Vpc struct {
	VpcId     *string `json:"vpcId,omitempty"`
	VpcName   *string `json:"vpcName,omitempty"`
	VpcRegion *string `json:"vpcRegion,omitempty"`
}
