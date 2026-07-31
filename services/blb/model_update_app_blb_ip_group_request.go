package blb

type UpdateAppBlbIpGroupRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
	IpGroupId   *string `json:"ipGroupId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Desc        *string `json:"desc,omitempty"`
}
