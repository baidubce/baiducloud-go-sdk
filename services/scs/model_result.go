package scs

type Result struct {
	TemplateId         *int32        `json:"templateId,omitempty"`
	TemplateShowId     *string       `json:"templateShowId,omitempty"`
	TemplateName       *string       `json:"templateName,omitempty"`
	ParametersNum      *int32        `json:"parametersNum,omitempty"`
	ClusterType        *string       `json:"clusterType,omitempty"`
	Engine             *string       `json:"engine,omitempty"`
	EngineVersion      *string       `json:"engineVersion,omitempty"`
	TemplateType       *int32        `json:"templateType,omitempty"`
	NeedReboot         *int32        `json:"needReboot,omitempty"`
	Comment            *string       `json:"comment,omitempty"`
	CreateTime         *string       `json:"createTime,omitempty"`
	UpdateTime         *string       `json:"updateTime,omitempty"`
	Parameters         []*Parameters `json:"parameters,omitempty"`
	CacheClusterShowId *string       `json:"cacheClusterShowId,omitempty"`
	CacheClusterName   *string       `json:"cacheClusterName,omitempty"`
	AvailabilityZone   *string       `json:"availabilityZone,omitempty"`
	Version            *int32        `json:"version,omitempty"`
	Status             *string       `json:"status,omitempty"`
	ApplyTime          *string       `json:"applyTime,omitempty"`
	ConfName           *string       `json:"confName,omitempty"`
	ConfDefault        *string       `json:"confDefault,omitempty"`
	ConfValue          *string       `json:"confValue,omitempty"`
	ConfType           *int32        `json:"confType,omitempty"`
	ConfRange          *string       `json:"confRange,omitempty"`
	ConfModule         *int32        `json:"confModule,omitempty"`
	ConfDesc           *string       `json:"confDesc,omitempty"`
	ConfRedisVersion   *string       `json:"confRedisVersion,omitempty"`
	ConfCacheVersion   *int32        `json:"confCacheVersion,omitempty"`
	LeaderName         *string       `json:"leaderName,omitempty"`
	LeaderShowId       *string       `json:"leaderShowId,omitempty"`
	LeaderRegion       *string       `json:"leaderRegion,omitempty"`
	GroupId            *string       `json:"groupId,omitempty"`
	GroupName          *string       `json:"groupName,omitempty"`
	GroupStatus        *string       `json:"groupStatus,omitempty"`
	ClusterNum         *int32        `json:"clusterNum,omitempty"`
	GroupCreateTime    *string       `json:"groupCreateTime,omitempty"`
	ForbidWrite        *int32        `json:"forbidWrite,omitempty"`
	GroupType          *string       `json:"groupType,omitempty"`
}
