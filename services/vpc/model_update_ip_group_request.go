package vpc

type UpdateIpGroupRequest struct {
	IpSetId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
