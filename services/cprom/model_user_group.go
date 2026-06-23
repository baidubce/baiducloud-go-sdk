package cprom

type UserGroup struct {
	GroupId     *string `json:"groupId,omitempty"`
	GroupName   *string `json:"groupName,omitempty"`
	Description *string `json:"description,omitempty"`
}
