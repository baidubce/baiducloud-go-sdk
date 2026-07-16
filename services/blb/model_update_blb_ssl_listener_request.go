package blb

type UpdateBlbSslListenerRequest struct {
	BlbId                      *string   `json:"-"`
	ClientToken                *string   `json:"-"`
	ListenerPort               *int32    `json:"-"`
	BackendPort                *int32    `json:"backendPort,omitempty"`
	Scheduler                  *string   `json:"scheduler,omitempty"`
	HealthCheckType            *string   `json:"healthCheckType,omitempty"`
	HealthCheckTimeoutInSecond *int32    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        *int32    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         *int32    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           *int32    `json:"healthyThreshold,omitempty"`
	CertIds                    []*string `json:"certIds,omitempty"`
	EncryptionType             *string   `json:"encryptionType,omitempty"`
	EncryptionProtocols        []*string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers             *string   `json:"appliedCiphers,omitempty"`
	DualAuth                   *bool     `json:"dualAuth,omitempty"`
	ClientCertIds              []*string `json:"clientCertIds,omitempty"`
}
