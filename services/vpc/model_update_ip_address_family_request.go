package vpc

type UpdateIpAddressFamilyRequest struct {
	IpGroupId   *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
