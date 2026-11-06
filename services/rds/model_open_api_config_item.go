package rds

type OpenApiConfigItem struct {
	Name          *string `json:"name,omitempty"`
	DefaultValue  *string `json:"defaultValue,omitempty"`
	Value         *string `json:"value,omitempty"`
	PendingValue  *string `json:"pendingValue,omitempty"`
	RdsType       *string `json:"type,omitempty"`
	Dynamic       *string `json:"dynamic,omitempty"`
	Modifiable    *string `json:"modifiable,omitempty"`
	AllowedValues *string `json:"allowedValues,omitempty"`
	Desc          *string `json:"desc,omitempty"`
	Etag          *string `json:"etag,omitempty"`
	ConfigType    *string `json:"configType,omitempty"`
	BestValue     *string `json:"bestValue,omitempty"`
	Attention     *string `json:"attention,omitempty"`
	IfCluster     *bool   `json:"ifCluster,omitempty"`
}
