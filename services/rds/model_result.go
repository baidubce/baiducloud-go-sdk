package rds

type Result struct {
	Id         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	DbType     *string `json:"dbType,omitempty"`
	DbVersion  *string `json:"dbVersion,omitempty"`
	Count      *int32  `json:"count,omitempty"`
	Desc       *string `json:"desc,omitempty"`
	Dynamic    *bool   `json:"dynamic,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	IdShow     *string `json:"idShow,omitempty"`
	Para       []*Para `json:"para,omitempty"`
	RdsType    *string `json:"type,omitempty"`
	Property   *string `json:"property,omitempty"`
}
