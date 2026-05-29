package blb

type CreateAppBlbIpGroupMemberRequest struct {
	BlbId       *string             `json:"-"`
	ClientToken *string             `json:"-"`
	IpGroupId   *string             `json:"ipGroupId,omitempty"`
	MemberList  []*AppIpGroupMember `json:"memberList,omitempty"`
}
