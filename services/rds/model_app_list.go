package rds

type AppList struct {
	AppId       *string `json:"appId,omitempty"`
	AppName     *string `json:"appName,omitempty"`
	Weight      *int32  `json:"weight,omitempty"`
	RoGroupId   *string `json:"roGroupId,omitempty"`
	SourceAppId *string `json:"sourceAppId,omitempty"`
	Status      *string `json:"status,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	UpdateTime  *string `json:"updateTime,omitempty"`
	AppStatus   *string `json:"appStatus,omitempty"`
	AppIdShort  *string `json:"appIdShort,omitempty"`
}
