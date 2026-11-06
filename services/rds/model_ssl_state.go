package rds

type SslState struct {
	InstanceId *string `json:"instanceId,omitempty"`
	SslStatus  *string `json:"sslStatus,omitempty"`
	SslAddress *string `json:"sslAddress,omitempty"`
	StartTime  *string `json:"startTime,omitempty"`
	EndTime    *string `json:"endTime,omitempty"`
}
