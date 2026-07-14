package as

type HealthCheckState struct {
	CheckId       *string        `json:"checkId,omitempty"`
	GroupId       *string        `json:"groupId,omitempty"`
	AccountId     *string        `json:"accountId,omitempty"`
	State         *string        `json:"state,omitempty"`
	CheckEntities []*CheckEntity `json:"checkEntities,omitempty"`
	CreateTime    *string        `json:"createTime,omitempty"`
}
