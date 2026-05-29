package blb

type DeleteAppBlbIpGroupProtocolRequest struct {
	BlbId               *string   `json:"-"`
	ClientToken         *string   `json:"-"`
	IpGroupId           *string   `json:"ipGroupId,omitempty"`
	BackendPolicyIdList []*string `json:"backendPolicyIdList,omitempty"`
}
