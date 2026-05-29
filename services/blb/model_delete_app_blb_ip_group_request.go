package blb

type DeleteAppBlbIpGroupRequest struct {
	BlbId       *string `json:"-"`
	ClientToken *string `json:"-"`
	IpGroupId   *string `json:"ipGroupId,omitempty"`
}
