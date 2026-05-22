package blb

type UpdateAppBlbHttpsListenerRequest struct {
	BlbId                 *string                 `json:"-"`
	ClientToken           *string                 `json:"-"`
	ListenerPort          *int32                  `json:"-"`
	Scheduler             *string                 `json:"scheduler,omitempty"`
	KeepSession           *bool                   `json:"keepSession,omitempty"`
	KeepSessionType       *string                 `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    *int32                  `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName *string                 `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         *bool                   `json:"xForwardedFor,omitempty"`
	XForwardedProto       *bool                   `json:"xForwardedProto,omitempty"`
	ServerTimeout         *int32                  `json:"serverTimeout,omitempty"`
	CertIds               []*string               `json:"certIds,omitempty"`
	EncryptionType        *string                 `json:"encryptionType,omitempty"`
	EncryptionProtocols   []*string               `json:"encryptionProtocols,omitempty"`
	AppliedCiphers        *string                 `json:"appliedCiphers,omitempty"`
	DualAuth              *bool                   `json:"dualAuth,omitempty"`
	ClientCertIds         []*string               `json:"clientCertIds,omitempty"`
	AdditionalCertDomains []*AdditionalCertDomain `json:"additionalCertDomains,omitempty"`
	Description           *string                 `json:"description,omitempty"`
}
