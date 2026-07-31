package blb

type CreateAppBlbIpGroupRequest struct {
	BlbId       *string                      `json:"-"`
	ClientToken *string                      `json:"-"`
	Name        *string                      `json:"name,omitempty"`
	Desc        *string                      `json:"desc,omitempty"`
	MemberList  []*AppIpGroupMemberForCreate `json:"memberList,omitempty"`
}
