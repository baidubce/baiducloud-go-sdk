package scs

type HotGroupModifyNameRequest struct {
	GroupId   *string `json:"-"`
	GroupName *string `json:"groupName,omitempty"`
}
