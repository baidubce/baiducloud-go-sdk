package rds

type ResourceGroup struct {
	Name       *string `json:"name,omitempty"`
	Extra      *string `json:"extra,omitempty"`
	ParentUuid *string `json:"parentUuid,omitempty"`
	GroupId    *string `json:"groupId,omitempty"`
	AccountId  *string `json:"accountId,omitempty"`
	UserId     *string `json:"userId,omitempty"`
	CreateTime *int64  `json:"createTime,omitempty"`
	UpdateTime *int64  `json:"updateTime,omitempty"`
	DeleteTime *int64  `json:"deleteTime,omitempty"`
	BindTime   *string `json:"bindTime,omitempty"`
}
