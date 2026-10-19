package rds

type ResourceInfo struct {
	Name       *string                 `json:"name,omitempty"`
	Extra      *string                 `json:"extra,omitempty"`
	ParentUuid *string                 `json:"parentUuid,omitempty"`
	GroupId    *string                 `json:"groupId,omitempty"`
	AccountId  *string                 `json:"accountId,omitempty"`
	UserId     *string                 `json:"userId,omitempty"`
	CreateTime *map[string]interface{} `json:"createTime,omitempty"`
	UpdateTime *map[string]interface{} `json:"updateTime,omitempty"`
	DeleteTime *map[string]interface{} `json:"deleteTime,omitempty"`
	BindTime   *map[string]interface{} `json:"bindTime,omitempty"`
}
