package privatezone

type CreateResolverRequest struct {
	ClientToken     *string    `json:"-"`
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	VpcId           *string    `json:"vpcId,omitempty"`
	VpcRegion       *string    `json:"vpcRegion,omitempty"`
	IpModels        []*IpModel `json:"ipModels,omitempty"`
	PrivatezoneType *string    `json:"type,omitempty"`
}
