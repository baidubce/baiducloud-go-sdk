package cprom

type MonitorInstance struct {
	InstanceId      *string `json:"instanceId,omitempty"`
	InstanceName    *string `json:"instanceName,omitempty"`
	InstanceSpec    *string `json:"instanceSpec,omitempty"`
	InstanceType    *string `json:"instanceType,omitempty"`
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	CreateTime      *string `json:"createTime,omitempty"`
	GrafanaId       *string `json:"grafanaId,omitempty"`
	GrafanaName     *string `json:"grafanaName,omitempty"`
	Status          *Status `json:"status,omitempty"`
}
