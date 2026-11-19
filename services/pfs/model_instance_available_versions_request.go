package pfs

type InstanceAvailableVersionsRequest struct {
	Action       *string `json:"-"`
	InstanceType *string `json:"instanceType,omitempty"`
	VpcId        *string `json:"vpcId,omitempty"`
}
