package rds

type TableGranularityQueryUserDetailsApiRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	User       *string `json:"user,omitempty"`
	Host       *string `json:"host,omitempty"`
}
