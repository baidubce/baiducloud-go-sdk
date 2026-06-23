package cprom

type CreateInstanceRequest struct {
	InstanceName         *string `json:"instanceName,omitempty"`
	InstanceType         *string `json:"instanceType,omitempty"`
	InstanceSpec         *string `json:"instanceSpec,omitempty"`
	RetentionPeriod      *string `json:"retentionPeriod,omitempty"`
	NeedGrafana          *bool   `json:"needGrafana,omitempty"`
	GrafanaName          *string `json:"grafanaName,omitempty"`
	GrafanaAdminPassword *string `json:"grafanaAdminPassword,omitempty"`
}
