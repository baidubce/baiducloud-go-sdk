package blb

type UpdateBlbHttpsListenerRequest struct {
	BlbId                      *string                    `json:"-"`
	ListenerPort               *int32                     `json:"-"`
	BackendPort                *int32                     `json:"backendPort,omitempty"`
	Scheduler                  *string                    `json:"scheduler,omitempty"`
	KeepSession                *bool                      `json:"keepSession,omitempty"`
	KeepSessionType            *string                    `json:"keepSessionType,omitempty"`
	KeepSessionDuration        *int32                     `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      *string                    `json:"keepSessionCookieName,omitempty"`
	XForwardFor                *bool                      `json:"xForwardFor,omitempty"`
	XForwardedProto            *bool                      `json:"xForwardedProto,omitempty"`
	AdditionalAttributes       *AdditionalAttributesModel `json:"additionalAttributes,omitempty"`
	HealthCheckType            *string                    `json:"healthCheckType,omitempty"`
	HealthCheckPort            *int32                     `json:"healthCheckPort,omitempty"`
	HealthCheckURI             *string                    `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond *int32                     `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        *int32                     `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         *int32                     `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           *int32                     `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    *string                    `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost            *string                    `json:"healthCheckHost,omitempty"`
	ServerTimeout              *int32                     `json:"serverTimeout,omitempty"`
	CertIds                    []*string                  `json:"certIds,omitempty"`
	AdditionalCertDomains      []*AdditionalCertDomain    `json:"additionalCertDomains,omitempty"`
	EncryptionType             *string                    `json:"encryptionType,omitempty"`
	EncryptionProtocols        []*string                  `json:"encryptionProtocols,omitempty"`
	AppliedCiphers             *string                    `json:"appliedCiphers,omitempty"`
}
