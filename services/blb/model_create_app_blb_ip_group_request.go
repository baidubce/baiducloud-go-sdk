package blb

type CreateAppBlbIpGroupRequest struct {
	BlbId                   *string                      `json:"-"`
	ClientToken             *string                      `json:"-"`
	Name                    *string                      `json:"name,omitempty"`
	Desc                    *string                      `json:"desc,omitempty"`
	PreserveClientIpEnabled *bool                        `json:"preserveClientIpEnabled,omitempty"`
	GroupTargetType         *string                      `json:"groupTargetType,omitempty"`
	MemberList              []*AppIpGroupMemberForCreate `json:"memberList,omitempty"`
}
