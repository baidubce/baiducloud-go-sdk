package cfs

type AccessGroupModel struct {
	AccessGroupName *string            `json:"accessGroupName,omitempty"`
	AccessRules     []*AccessRuleModel `json:"accessRules,omitempty"`
	CreateTime      *string            `json:"createTime,omitempty"`
	Description     *string            `json:"description,omitempty"`
	FsCount         *int32             `json:"fsCount,omitempty"`
}
