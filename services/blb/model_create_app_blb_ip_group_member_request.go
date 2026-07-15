package blb

type CreateAppBlbIpGroupMemberRequest struct {
	BlbId       *string                      `json:"-"`
	ClientToken *string                      `json:"-"`
	IpGroupId   *string                      `json:"ipGroupId,omitempty"`
	MemberList  []*AppIpGroupMemberForCreate `json:"memberList,omitempty"`
}
