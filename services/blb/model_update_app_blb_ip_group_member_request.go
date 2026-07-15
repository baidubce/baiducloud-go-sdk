package blb

type UpdateAppBlbIpGroupMemberRequest struct {
	BlbId       *string                      `json:"-"`
	ClientToken *string                      `json:"-"`
	IpGroupId   *string                      `json:"ipGroupId,omitempty"`
	MemberList  []*AppIpGroupMemberForUpdate `json:"memberList,omitempty"`
}
