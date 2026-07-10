package pfs

type CreatePfsRequest struct {
	Name         *string `json:"name,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	Capacity     *int32  `json:"capacity,omitempty"`
	Zone         *string `json:"zone,omitempty"`
	SubnetId     *string `json:"subnetId,omitempty"`
	Description  *string `json:"description,omitempty"`
	Tags         []*Tag  `json:"tags,omitempty"`
}
