package blb

type DeleteAppBlbIpGroupMemberRequest struct {
	BlbId        *string   `json:"-"`
	ClientToken  *string   `json:"-"`
	IpGroupId    *string   `json:"ipGroupId,omitempty"`
	MemberIdList []*string `json:"memberIdList,omitempty"`
}
